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
	id, isRight := service.CheckLogin(msg.Username, msg.Password)
	if isRight {
		token, err := service.GenerateToken(msg.Username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "登录失败"})
		}
		if err := service.SaveToken(token, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "登录失败"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": gin.H{
			"token": token,
		}})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误", "message": "登录失败"})
	}
}

func GetAdminByAuth(c *gin.Context) {
	id, exists := c.Get("admin_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"admin": gin.H{
		"id":       id,
		"username": username,
	}})
}
