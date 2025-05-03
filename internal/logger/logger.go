package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	Log  *logrus.Logger
	once sync.Once
)

func Init() {
	once.Do(func() {
		Log = logrus.New()
		Log.SetFormatter(&logrus.JSONFormatter{})
		Log.SetOutput(os.Stdout)
		Log.SetLevel(logrus.InfoLevel)
	})
}
