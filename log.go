package webhookLog

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

type Logger interface {
	SetLevel(level LogLevel)
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type DefaultLogger struct {
	name  string
	whid  string
	key   string
	url   string
	level LogLevel
}

func NewDefaultLogger(name string, key string) *DefaultLogger {
	return &DefaultLogger{
		level: Info,
		name:  name,
		key:   key,
		url:   "https://discord.com/api/webhooks/" + key,
	}
}

func (l *DefaultLogger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *DefaultLogger) Debugf(msg string, args ...interface{}) {
	if l.level <= Debug {
		l.log(Debug, msg, args...)
	}
}

func (l *DefaultLogger) Infof(msg string, args ...interface{}) {
	if l.level <= Info {
		l.log(Info, msg, args...)
	}
}

func (l *DefaultLogger) Warnf(msg string, args ...interface{}) {
	if l.level <= Warn {
		l.log(Warn, msg, args...)
	}
}

func (l *DefaultLogger) Errorf(msg string, args ...interface{}) {
	if l.level <= Error {
		l.log(Error, msg, args...)
	}
}

func (l *DefaultLogger) Fatalf(msg string, args ...interface{}) {
	if l.level <= Fatal {
		l.log(Fatal, msg, args...)
		os.Exit(1)
	}
}

func (l *DefaultLogger) Debug(msg string) {
	l.Debugf(msg)
}

func (l *DefaultLogger) Info(msg string) {
	l.Infof(msg)
}

func (l *DefaultLogger) Warn(msg string) {
	l.Warnf(msg)
}

func (l *DefaultLogger) Error(msg string) {
	l.Errorf(msg)
}

func (l *DefaultLogger) Fatal(msg string) {
	l.Fatalf(msg)
}

func (l *DefaultLogger) log(level LogLevel, msg string, args ...interface{}) {
	formattedMsg := fmt.Sprintf("[%s] [%s] %s \\n", time.Now().Format(time.RFC3339), levelToString(level), fmt.Sprintf(msg, args...))
	l.sendMessage(formattedMsg)
}

func levelToString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return ""
}
