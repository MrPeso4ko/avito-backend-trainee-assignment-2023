package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

const defaultLogLevel = logrus.InfoLevel

var entry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func init() {
	logLevel := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = defaultLogLevel
	}
	logger := logrus.New()
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	if os.Getenv("MODE") == "prod" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	entry = logrus.NewEntry(logger)
}
