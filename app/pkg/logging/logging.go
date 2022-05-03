package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// writerHook is a hook that writes logs of specified LogLevels to specified Writer
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) LWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func (l *Logger) LWithFields(fields map[string]interface{}) *Logger {
	return &Logger{l.WithFields(fields)}
}

func Init(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatalln(err)
	}

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	l.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

	l.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrusLevel)

	e = logrus.NewEntry(l)
}
