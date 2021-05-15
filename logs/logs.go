package logs

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type logrusToEchoLogger struct {
	*logrus.Logger
}

func LogrusAdapter() echo.Logger {
	return &logrusToEchoLogger{Logger: logrus.StandardLogger()}
}

func (l *logrusToEchoLogger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.FatalLevel:
		return log.ERROR
	case logrus.PanicLevel:
		return log.ERROR
	default:
		return log.OFF
	}
}

func (l *logrusToEchoLogger) SetLevel(v log.Lvl) {
	switch v {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.InfoLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.PanicLevel)
	}
}

func (l *logrusToEchoLogger) Output() io.Writer {
	return l.Logger.Out
}

func (l *logrusToEchoLogger) Prefix() string {
	return ""
}

func (l *logrusToEchoLogger) SetPrefix(p string) {
	//noop
}

func (l *logrusToEchoLogger) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	l.Print(string(b))
}

func (l *logrusToEchoLogger) SetHeader(h string) {
	//noop
}

func (l *logrusToEchoLogger) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Info(string(b))
}

func (l *logrusToEchoLogger) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Warn(string(b))
}

func (l *logrusToEchoLogger) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Debug(string(b))
}

func (l *logrusToEchoLogger) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Error(string(b))
}

func (l *logrusToEchoLogger) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Fatal(string(b))
}

func (l *logrusToEchoLogger) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	l.Panic(string(b))
}
