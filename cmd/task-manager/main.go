package main

import (
	"os"

	"github.com/helloamitkr/go-task-manager-app/internal/logger"
	"github.com/helloamitkr/go-task-manager-app/internal/server"
)

func main() {
	if os.Getenv("GIN_MODE") != "release" {
		os.Setenv("GIN_MODE", "release")
	}
	logger.Init()
	srv := server.NewServer()
	srv.Run("8000")
}
