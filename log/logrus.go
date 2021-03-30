package log

import "github.com/sirupsen/logrus"

func NewLogrusLogger() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.TraceLevel)
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	return &loggerWrapper{logrusLogger: logrusLogger}
}

type loggerWrapper struct {
	logrusLogger *logrus.Logger
}

func (logger *loggerWrapper) Info(args ...interface{}) {
	logger.logrusLogger.Info(args)
}

func (logger *loggerWrapper) Debug(args ...interface{}) {
	logger.logrusLogger.Debug(args)
}

func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
	logger.logrusLogger.Errorf(format, args)
}

func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
	logger.logrusLogger.Fatalf(format, args)
}

func (logger *loggerWrapper) Fatal(args ...interface{}) {
	logger.logrusLogger.Fatal(args)
}

func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
	logger.logrusLogger.Infof(format, args)
}

func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
	logger.logrusLogger.Warnf(format, args)
}

func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
	logger.logrusLogger.Debugf(format, args)
}

func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
	logger.logrusLogger.Infof(format, args)
}

func (logger *loggerWrapper) Println(args ...interface{}) {
	logger.logrusLogger.Info(args, "\n")
}
