package router

import (
	"EmqxBackEnd/handlers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.POST("/empx", handlers.ReceiveEmpx)
	return r
}
