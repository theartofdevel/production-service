package logging

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type logger struct {
	*logrus.Logger
}

type Logger interface {
	SetLevel(level logrus.Level)
	GetLevel() logrus.Level
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
	WithContext(ctx context.Context) *logrus.Entry
	WithTime(t time.Time) *logrus.Entry
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

func GetLogger(ctx context.Context) Logger {
	return loggerFromContext(ctx)
}

func NewLogger() Logger {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.InfoLevel)
	logrusLogger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	logrusLogger.SetOutput(os.Stdout)
	// TODO не работает, починить
	logrusLogger.SetReportCaller(true)

	return &logger{
		Logger: logrusLogger,
	}
}

func (l *logger) SetLevel(level logrus.Level) {
	l.Logger.SetLevel(level)
}

func (l *logger) GetLevel() logrus.Level {
	return l.Logger.GetLevel()
}

func (l *logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.Logger.WithField(key, value)
}

func (l *logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

func (l *logger) WithError(err error) *logrus.Entry {
	return l.Logger.WithError(err)
}

func (l *logger) WithContext(ctx context.Context) *logrus.Entry {
	return l.Logger.WithContext(ctx)
}

func (l *logger) WithTime(t time.Time) *logrus.Entry {
	return l.Logger.WithTime(t)
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.Logger.Tracef(format, args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *logger) Warningf(format string, args ...interface{}) {
	l.Logger.Warningf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.Logger.Panicf(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	l.Logger.Traceln(args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.Logger.Debugln(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.Logger.Infoln(args...)
}

func (l *logger) Print(args ...interface{}) {
	l.Logger.Println(args...)
}

func (l *logger) Warning(args ...interface{}) {
	l.Logger.Warningln(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.Logger.Errorln(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Logger.Fatalln(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.Logger.Panicln(args...)
}
