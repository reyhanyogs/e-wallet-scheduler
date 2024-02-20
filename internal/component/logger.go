package component

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

var Log = initializeLogger()

func initializeLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	file, err := os.OpenFile("logs/e-wallet-scheduler.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}

	log.Out = file

	log.SetFormatter(&ecslogrus.Formatter{})
	return log
}
