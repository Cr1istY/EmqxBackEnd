package handlers

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var msg models.EmpxAdmin
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if service.CheckLogin(msg.Username, msg.Password) {
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误", "message": "登录失败"})
	}
}
