package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var inited = false

func GetLogger(module, place string) *log.Entry {
	if !inited {
		Init()
		inited = true
	}
	logger := log.New().WithFields(log.Fields{
		"module": module,
		"place": place,
	})
	return logger
}

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}