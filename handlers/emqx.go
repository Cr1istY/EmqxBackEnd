package handlers

import (
	"EmqxBackEnd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReceiveEmpx(c *gin.Context) {
	var msg models.EmpxMessage
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
