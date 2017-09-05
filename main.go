package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

func main() {
	// So that the order of the log and our Prints is correct
	logrus.SetOutput(os.Stdout)

	// Uncomment the next line to enable debug output - default is info and higher
	//logrus.SetLevel(logrus.DebugLevel)

	lvl, err := getLevel(logrus.StandardLogger())
	if err != nil {
		fmt.Printf("failed to get logger level: %s\n\n", err)
	} else {
		fmt.Printf("logger level: %s\n\n", lvl)
	}

	test("debug", logrus.StandardLogger())
	fmt.Println()
	test("info", logrus.StandardLogger())
	fmt.Println()
	test("warn", logrus.StandardLogger())
}

func test(level string, logger logrus.FieldLogger) {
	fmt.Println("Level is", level)
	fmt.Printf("%v\n", reflect.TypeOf(logger))
	switch level {
	case "debug":
		logger.Debug("test")
	case "info":
		logger.Info("test")
	case "warn":
		logger.Warn("test")
	}
}

func getLevel(logger logrus.FieldLogger) (logrus.Level, error) {
	switch l := logger.(type) {
	case *logrus.Logger:
		return l.Level, nil
	case *logrus.Entry:
		return l.Level, nil
	default:
		return logrus.PanicLevel, fmt.Errorf("unknown logger type")
	}
}
