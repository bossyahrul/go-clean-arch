package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func NewLogrusLogger() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.TraceLevel)
	logrusLogger.SetReportCaller(true)
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	return &LoggerWrapper{LogrusLogger: logrusLogger}
}

type LoggerWrapper struct {
	LogrusLogger *logrus.Logger
}

func (logger *LoggerWrapper) Info(args ...interface{}) {
	logger.LogrusLogger.Info(args)
}

func (logger *LoggerWrapper) Debug(args ...interface{}) {
	logger.LogrusLogger.Debug(args)
}

func (logger *LoggerWrapper) Errorf(format string, args ...interface{}) {
	logger.LogrusLogger.Errorf(format, args)
}

func (logger *LoggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.LogrusLogger.Fatalf(format, args)
}

func (logger *LoggerWrapper) Fatal(args ...interface{}) {
	logger.LogrusLogger.Fatal(args)
}

func (logger *LoggerWrapper) Infof(format string, args ...interface{}) {
	logger.LogrusLogger.Infof(format, args)
}

func (logger *LoggerWrapper) Warnf(format string, args ...interface{}) {
	logger.LogrusLogger.Warnf(format, args)
}

func (logger *LoggerWrapper) Debugf(format string, args ...interface{}) {
	logger.LogrusLogger.Debugf(format, args)
}

func (logger *LoggerWrapper) Printf(format string, args ...interface{}) {
	logger.LogrusLogger.Infof(format, args)
}

func (logger *LoggerWrapper) Println(args ...interface{}) {
	logger.LogrusLogger.Info(args, "\n")
}
