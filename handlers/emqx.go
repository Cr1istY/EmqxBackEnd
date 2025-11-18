package handlers

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
	"EmqxBackEnd/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReceiveEmpx(c *gin.Context) {
	var msg models.EmpxMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ProcessEmpxMessage(&msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "saved"})
}

func GetMessages(c *gin.Context) {
	var msg models.GetMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 取出当前用户的token，与用户的id的token进行对比
	token := c.GetHeader("Authorization")
	userToken, err := repository.GetToken(msg.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get token"})
		return
	}
	if token != userToken {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	messages, err := repository.GetMessages(msg.Type, msg.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}
	c.JSON(http.StatusOK, messages)
}
