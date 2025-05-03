package server

import (
	"github.com/gin-gonic/gin"
	"github.com/helloamitkr/go-task-manager-app/internal/logger"
	"github.com/helloamitkr/go-task-manager-app/internal/router"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: router.InitilizeRouter(),
	}
}

func (s *Server) Run(port string) {
	logger.Log.Infof("Start the server at port %s", port)
	if err := s.router.Run(":" + port); err != nil {
		logger.Log.Fatalf("Start the server at port %s", port)
	}
}
