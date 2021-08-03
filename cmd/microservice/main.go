package main

import (
	"fmt"
	"github.com/haf80/microservice/pkg/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.NewLogrusLogger(logrus.InfoLevel, &logrus.JSONFormatter{})
	logger.Info(map[string]interface{}{"name": "soroush"}, "test message")
	var a map[string]interface{}
	b := logrus.Fields(a)
	fmt.Println(b)
}
