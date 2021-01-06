package logger

import "github.com/sirupsen/logrus"

var (
	Logger *logrus.Logger
)

func InitLogger() *logrus.Logger {
	Logger = logrus.New()
	return Logger
}

func GetLogger() *logrus.Logger {
	return Logger
}
