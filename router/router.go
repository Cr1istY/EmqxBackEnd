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
	r.POST("/empx", handlers.ReceiveEmpx)
	r.POST("/admin/login", handlers.Login)
	protected := r.Group("/emqx/dashboard")
	protected.Use(middleware.AuthMiddlewareWithCache())
	{
		protected.GET("/:type", handlers.GetMessages)
	}

	return r
}
