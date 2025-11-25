package handlers

import (
	"EmqxBackEnd/models"
	"EmqxBackEnd/repository"
	"EmqxBackEnd/service"
	"EmqxBackEnd/state"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ReceiveEmpx(c *gin.Context) {
	var emqxMsg models.EMQXMessagePublish
	if err := c.ShouldBindJSON(&emqxMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(emqxMsg)
	var rawMsg models.RawEmpxMessage
	if err := json.Unmarshal([]byte(emqxMsg.Payload), &rawMsg); err != nil {
		log.Println("Error unmarshalling payload:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload 解析失败"})
		return
	}
	log.Println(rawMsg)
	var msg models.EmpxMessage
	msg.Value = rawMsg.Value
	msg.NodeID = rawMsg.NodeID
	msg.Type = rawMsg.Type
	msg.TS = time.Now()

	if err := service.ProcessEmpxMessage(&msg); err != nil {
		log.Println("Error saving message:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	if msg.Type == 4 {
		ppm, err := strconv.Atoi(msg.Value)
		if err != nil {
			log.Println("Error converting ppm to int:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "ppm 解析失败"})
			return
		}
		if ppm >= 2100 {
			// 进入危险值
			state.SetCache("ppm", 3) // 打开蜂鸣器
		} else {
			state.SetCache("ppm", 4) // 关闭蜂鸣器
		}
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

	var messages []models.EmpxMessage

	if messageType == "3" || messageType == "4" {
		var messages3 []models.EmpxMessage
		messages3, err = repository.GetMessages(3, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
			return
		}
		var messages4 []models.EmpxMessage
		messages4, err = repository.GetMessages(4, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
			return
		}
		messages = append(messages3, messages4...)
	} else {
		messages, err = repository.GetMessages(int(messageTypeId), userId)
	}
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}
	if len(messages) == 0 {
		c.JSON(http.StatusOK, gin.H{"messages": []models.EmpxMessage{}})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func Empx(c *gin.Context) {
	log.Println("Empx", c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
