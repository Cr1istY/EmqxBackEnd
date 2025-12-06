package main

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/handlers"
	"EmqxBackEnd/jobs"
	"EmqxBackEnd/mqtt"
	"EmqxBackEnd/router"
	"EmqxBackEnd/state"
	"EmqxBackEnd/task"
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mqttBroker := "mqtt://172.20.10.5:1883"
	mqttUser := ""
	mqttPass := ""
	if err := mqtt.InitClient(mqttBroker, "cron_task_client", mqttUser, mqttPass); err != nil {
		log.Fatalf("MQTT初始化失败: %v", err)
	}
	defer mqtt.Close()

	db, err := database.Init()
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	state.SetCache("ppm", 4)

	taskMgr := task.NewManager(db)
	taskMgr.RegisterTask("温度传感器数据", jobs.GetTem)
	taskMgr.RegisterTask("获取气体ppm值", jobs.GetPPM)
	taskMgr.RegisterTask("获取空气湿度", jobs.GetMoisture)
	taskMgr.RegisterTask("获取红外传感器数据", jobs.GetInfrared)

	if err := taskMgr.LoadTasksFromDB(); err != nil {
		log.Printf("⚠️ 加载任务失败: %v", err)
	}

	// 6. 启动Cron调度器
	taskMgr.StartCron()
	defer taskMgr.StopCron()

	// 7. 将任务管理器注入到handler层
	handlers.SetTaskManager(taskMgr)

	r := router.Setup()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("\n收到退出信号，正在关闭服务...")

		// 5秒超时
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//停止定时任务
		taskMgr.StopCron()

		// 关闭数据库连接
		_ = db.Close()

		// 断开MQTT连接
		mqtt.Close()

		log.Println("所有资源已释放，服务已停止")
		os.Exit(0)
	}()

	log.Println("🚀 MQTT定时任务服务启动在 :8080")

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
