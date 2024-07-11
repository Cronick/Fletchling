package util

import (
	"encoding/json"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

var panicLogger *logrus.Logger

func SetPanicLogger(logger *logrus.Logger) {
	panicLogger = logger
}

func HandlePanic() {
	logger := panicLogger
	if logger != nil {
		if rerr := recover(); rerr != nil {
			stack := debug.Stack()
			jsonStack, err := json.Marshal(strings.Split(string(stack), "\n"))
			if err != nil {
				jsonStack = stack
			}
			logger.Errorf("Ruh roh. A panic has occurred: %v, trace: %s", rerr, string(jsonStack))
			panic(rerr)
		}
	}
}
