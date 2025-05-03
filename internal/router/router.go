package router

import (
	"github.com/gin-gonic/gin"
	"github.com/helloamitkr/go-task-manager-app/internal/handlers"
)

func InitilizeRouter() *gin.Engine {
	router := gin.New()
	healthrouter := router.Group("/health")
	{
		healthrouter.GET("/", handlers.HealthOk)
	}
	taskRouter := router.Group("/users")
	{
		taskRouter.POST("/", handlers.CreateNewTask)
	}
	return router
}
