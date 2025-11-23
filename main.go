package main

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/router"
	"database/sql"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func init() {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).    // MQTT服务器地址
		SetClientID(clientID) // 客户端ID

	mqttClient := mqtt.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT连接失败: %v", token.Error())
	}
}

func main() {
	database.Init()
	defer func(DB *sql.DB) {
		_ = DB.Close()
	}(database.DB)
	r := router.Setup()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
