package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// InitLogging initialize the logging module
func InitLogging() {
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	if os.Getenv("ENV") == "PROD" {
		file, err := os.OpenFile(os.Getenv("LOG_DIR")+"/yourapp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Out = file
		} else {
			logger.Info("Failed to log to file, using default stderr")
		}
	}
}

// GetLogger returns a logrus logger
func GetLogger() *logrus.Logger {
	return logger
}

// GetLoggerWithAction returns a logrus logger with preconfigured field "action": input
func GetLoggerWithAction(value string) *logrus.Entry {
	return logger.WithField("action", value)
}
