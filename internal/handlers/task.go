package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewTask(c *gin.Context) {
	c.String(http.StatusOK, "Post is called")
}
func GetAllTask(c *gin.Context) {
	c.String(http.StatusOK, "Get is called")
}
