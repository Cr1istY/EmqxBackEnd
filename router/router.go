package router

import (
	"EmqxBackEnd/handlers"
	"EmqxBackEnd/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.POST("/empx", handlers.Empx)
	r.POST("/empx/saveMessage", handlers.ReceiveEmpx)
	r.POST("/admin/login", handlers.Login)
	protected := r.Group("")
	protected.Use(middleware.AuthMiddlewareWithCache())
	{
		protected.GET("/admin/getinfo", handlers.GetAdminByAuth)
		// protected.GET("/empx/getNodeMessage", handlers.GetMessages)
		protected.GET("/empx/getMessage/:type", handlers.GetMessages)
		// protected.POST("/empx/getMessageByDaily", handlers.GetMessagesByDaily)
		protected.GET("empx/openTheDoor/:nodeId", handlers.OpenTheDoor)
		protected.GET("empx/openTheDoor/:nodeId", handlers.CloseTheDoor)
		protected.POST("/admin/register", handlers.Register)
		protected.POST("/admin/saveNode", handlers.SaveNode)
		protected.POST("/admin/changeUserStatus", handlers.ChangeUserStatus)
		protected.GET("/admin/getAllUser", handlers.GetAllUsers)
		protected.GET("/admin/getAllNode", handlers.GetAllNodeByUserId)
	}
	taskGroup := protected.Group("/task")
	{
		taskGroup.GET("", handlers.GetTasksHandler)                      // 获取任务列表
		taskGroup.PUT("/:name/cron", handlers.UpdateTaskCronHandler)     // 更新Cron表达式
		taskGroup.PUT("/:name/status", handlers.UpdateTaskStatusHandler) // 启用/禁用任务
	}
	return r
}
