package handlers

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
	"EmqxBackEnd/service"
	"net/http"
	"strconv"

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
	// 取出当前用户的token，与用户的id的token进行对比
	token := c.GetHeader("Authorization")
	messageType := c.Param("type")
	userId, err := repository.GetUserIdByToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get token"})
		return
	}
	messageTypeId, err := strconv.ParseInt(messageType, 10, 32)
	if err != nil {
		// Handle conversion error
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid type parameter"})
		return
	}
	messages, err := repository.GetMessages(int(messageTypeId), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}
	if len(messages) == 0 {
		c.JSON(http.StatusOK, gin.H{"messages": []models.EmpxMessage{}})
		return
	}
	c.JSON(http.StatusOK, messages)
}
