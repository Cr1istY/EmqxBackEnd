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
	r.POST("/empx/saveMessage", handlers.ReceiveEmpx)
	r.POST("/admin/login", handlers.Login)
	protected := r.Group("")
	protected.Use(middleware.AuthMiddlewareWithCache())
	{
		protected.GET("/admin/getinfo", handlers.GetAdminByAuth)
		protected.POST("/empx/getNodeMessage", handlers.GetMessages)
		protected.POST("/admin/register", handlers.Register)
	}
	return r
}
