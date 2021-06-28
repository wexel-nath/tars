package log

import (
	"fmt"
	"os"

	"tars/pkg/config"
)

var (
	levelMap = map[int]string {
		LevelDebug: "DEBUG",
		LevelInfo:  "INFO",
		LevelWarn:  "WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
	}
)

const (
	LevelDebug = 10
	LevelInfo  = 20
	LevelWarn  = 30
	LevelError = 40
	LevelFatal = 50
)

func getLevelString(level int) string {
	levelString, ok := levelMap[level]
	if ok {
		return levelString
	}

	return "INFO"
}

func Debug(format string, a ...interface{}) {
	log(LevelDebug, format, a...)
}

func Info(format string, a ...interface{}) {
	log(LevelInfo, format, a...)
}

func Warn(err error, a ...interface{}) {
	log(LevelWarn, err.Error(), a...)
}

func Error(err error, a ...interface{}) {
	log(LevelError, err.Error(), a...)
}

func Fatal(err error, a ...interface{}) {
	log(LevelFatal, err.Error(), a...)
	os.Exit(1)
}

func log(level int, format string, a ...interface{}) {
	if level >= config.Get().LogLevel {
		fmt.Println("[" + getLevelString(level) + "] " + fmt.Sprintf(format, a...))
	}
}
