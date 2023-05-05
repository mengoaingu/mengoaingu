package logger

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type LogrusHandler struct {
	logger *log.Logger
}

func NewLogrusHandler(logger *log.Logger) *LogrusHandler {
	return &LogrusHandler{
		logger: logger,
	}
}

func ConvertLogLevel(level string) log.Level {
	var l log.Level

	switch strings.ToLower(level) {
	case "error":
		l = log.ErrorLevel
	case "warm":
		l = log.WarnLevel
	case "info":
		l = log.InfoLevel
	case "debug":
		l = log.DebugLevel
	default:
		l = log.InfoLevel
	}

	return l
}
