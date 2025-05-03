package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helloamitkr/go-task-manager-app/internal/logger"
)

func HealthOk(c *gin.Context) {
	logger.Log.Info("Health endpoint:")
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}
