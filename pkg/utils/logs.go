package utils

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	// default is info level
	Logger.SetLevel(logrus.InfoLevel)
}

func SetLevel(level logrus.Level) {
	Logger.SetLevel(level)
}
