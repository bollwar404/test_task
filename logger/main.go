package logger

import (
	log "github.com/sirupsen/logrus"
)

func GetLogger(module, place string) *log.Entry {
	return log.WithFields(log.Fields{
		"module": module,
		"place": place,
	})
}