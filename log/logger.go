package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.Formatter = &logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	}
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.DebugLevel) // shut logger up during tests
}

func Warnf(fmt string, args ...interface{}) {
	Logger.Warnf(fmt, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Infof(fmt string, args ...interface{}) {
	Logger.Infof(fmt, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Debugf(fmt string, args ...interface{}) {
	Logger.Debugf(fmt, args...)
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Errorf(fmt string, args ...interface{}) {
	Logger.Errorf(fmt, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Fatalf(fmt string, args ...interface{}) {
	Logger.Fatalf(fmt, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
