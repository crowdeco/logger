package logger

import (
	"os"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type logger struct {
	logger *logrus.Logger
	debug  bool
}

func NewLogger() logger {
	debug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	pool := logrus.New()
	pool.SetFormatter(&logrus.JSONFormatter{})

	return logger{
		logger: pool,
		debug:  debug,
	}
}

func (l logger) Trace(message string) {
	if l.debug {
		var file string
		var line int
		var caller string

		pc, file, line, ok := runtime.Caller(1)
		detail := runtime.FuncForPC(pc)
		if ok || detail != nil {
			caller = detail.Name()
		}

		go l.logger.WithFields(l.fields(caller, file, line)).Trace(message)
	}
}

func (l logger) Debug(message string) {
	if l.debug {
		var file string
		var line int
		var caller string

		pc, file, line, ok := runtime.Caller(1)
		detail := runtime.FuncForPC(pc)
		if ok || detail != nil {
			caller = detail.Name()
		}

		go l.logger.WithFields(l.fields(caller, file, line)).Debug(message)
	}
}

func (l logger) Info(message string) {
	if l.debug {
		var file string
		var line int
		var caller string

		pc, file, line, ok := runtime.Caller(1)
		detail := runtime.FuncForPC(pc)
		if ok || detail != nil {
			caller = detail.Name()
		}

		go l.logger.WithFields(l.fields(caller, file, line)).Info(message)
	}
}

func (l logger) Warning(message string) {
	if l.debug {
		var file string
		var line int
		var caller string

		pc, file, line, ok := runtime.Caller(1)
		detail := runtime.FuncForPC(pc)
		if ok || detail != nil {
			caller = detail.Name()
		}

		go l.logger.WithFields(l.fields(caller, file, line)).Warning(message)
	}
}

func (l logger) Error(message string) {
	var file string
	var line int
	var caller string

	pc, file, line, ok := runtime.Caller(1)
	detail := runtime.FuncForPC(pc)
	if ok || detail != nil {
		caller = detail.Name()
	}

	go l.logger.WithFields(l.fields(caller, file, line)).Error(message)
}

func (l logger) Fatal(message string) {
	var file string
	var line int
	var caller string

	pc, file, line, ok := runtime.Caller(1)
	detail := runtime.FuncForPC(pc)
	if ok || detail != nil {
		caller = detail.Name()
	}

	go l.logger.WithFields(l.fields(caller, file, line)).Fatal(message)
}

func (l logger) Panic(message string) {
	var file string
	var line int
	var caller string

	pc, file, line, ok := runtime.Caller(1)
	detail := runtime.FuncForPC(pc)
	if ok || detail != nil {
		caller = detail.Name()
	}

	go l.logger.WithFields(l.fields(caller, file, line)).Panic(message)
}

func (l logger) fields(caller string, file string, line int) logrus.Fields {
	return logrus.Fields{
		"Debug":  l.debug,
		"Caller": caller,
		"File":   file,
		"Line":   line,
	}
}
