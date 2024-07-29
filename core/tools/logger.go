package tools

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	return log
}
