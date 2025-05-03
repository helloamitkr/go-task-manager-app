package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/helloamitkr/go-task-manager-app/internal/logger"
)

func LoggerWithLogrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		start := time.Now()

		// Continue the request chain
		c.Next()

		// Log the response status and duration
		duration := time.Since(start)
		logger.Log.Infof("Request: %s %s | Status: %d | Duration: %v", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}

func InitilizeRouter() *gin.Engine {
	router := gin.New()
	return router
}
