package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	With(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}

var lg Logger

func GetLogger() Logger {
	return lg
}

type Fields map[string]interface{}

type logger struct {
	entry *logrus.Entry
}

type Option struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

func NewLogger(w io.Writer, opt ...*Option) Logger {
	l := logrus.New()
	l.Out = w
	switch opt[0].Format {
	case "json":
		l.Formatter = &logrus.JSONFormatter{}
	default:
		l.Formatter = &logrus.TextFormatter{}
	}
	lvl, err := logrus.ParseLevel(opt[0].Level)
	if err != nil {
		panic(err)
	}
	l.SetLevel(lvl)
	return logger{logrus.NewEntry(l)}
}

func NewSingleLogger(w io.Writer, opt ...*Option) {
	lg = NewLogger(w, opt...)
}

func (l logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l logger) With(key string, value interface{}) Logger {
	return logger{l.entry.WithField(key, value)}
}

func (l logger) WithFields(fields Fields) Logger {
	return logger{l.entry.WithFields(logrus.Fields(fields))}
}

func (l logger) WithError(err error) Logger {
	return logger{l.entry.WithError(err)}
}
