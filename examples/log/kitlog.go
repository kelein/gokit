package main

import (
	"log"

	"github.com/kelein/gokit/pkg/logger"
)

func main() {
	opts := logger.Option{Code: logger.LOGRUS}
	logger, err := logger.Build(&opts)
	if err != nil {
		log.Fatal("Build logger failed")
	}
	logger.Info("Test Logger Info")
	logger.Error("Test Logger Info")
	logger.Warn("Test Logger Info")
	logger.Debug("Test Logger Info")
}
