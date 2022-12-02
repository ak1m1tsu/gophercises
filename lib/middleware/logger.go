package middleware

import (
	"log"
	"os"
	"time"
)

const (
	LoggerDefaultDateFormat = time.RFC3339
)

type ILogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type Logger struct {
	ILogger
	dateFormat string
}

func NewLogger() *Logger {
	return &Logger{
		ILogger:    log.New(os.Stdout, "[ak1m1tsu] ", 0),
		dateFormat: LoggerDefaultDateFormat,
	}
}

func (l *Logger) SetDateFormat(format string) {
	l.dateFormat = format
}
