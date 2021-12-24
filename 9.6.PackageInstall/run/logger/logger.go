package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	timeFormat string
	debug      bool
}

// Get instance of logger
func Instance(timeFormat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeFormat,
		debug:      debug,
	}
}

// Write log in instance
func (l Logger) Log(level string, s string) {
	switch level {
	case "info", "warning":
		if !l.debug {
			return
		}
	}
	fmt.Printf("[%s] %s %s\n", level, time.Now().Format(l.timeFormat), s)
}
