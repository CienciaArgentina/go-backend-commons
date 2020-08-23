package clog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

/*
 (C)iencia Argentina (Log)ger
*/

var (
	Logger *logrus.Logger

	TraceLevel = logrus.TraceLevel
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	FatalLevel = logrus.FatalLevel
	PanicLevel = logrus.PanicLevel
)

const (
	Level   = "level"
	Type    = "type"
	Subtype = "subtype"
)

func init() {
	Logger = &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Formatter: &logrus.JSONFormatter{},
		Level:     InfoLevel,
	}
}

func SetLogLevel(level logrus.Level) {
	if level.String() != "" {
		Logger.Level = level
	}
}

func Trace(message, logType string, tags map[string]string) {
	if Logger.Level >= logrus.TraceLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = TraceLevel.String()
		tags[Type] = logType
		Logger.WithFields(formatFields(tags)).Trace(message)
	}
}

func Debug(message, logType string, tags map[string]string) {
	if Logger.Level >= logrus.DebugLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = DebugLevel.String()
		tags[Type] = logType
		Logger.WithFields(formatFields(tags)).Debug(message)
	}
}

func Info(message, logType string, tags map[string]string) {
	if Logger.Level >= logrus.InfoLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = InfoLevel.String()
		tags[Type] = logType
		Logger.WithFields(formatFields(tags)).Info(message)
	}
}

func Warn(message, logType string, tags map[string]string) {
	if Logger.Level >= logrus.WarnLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = WarnLevel.String()
		tags[Type] = logType
		Logger.WithFields(formatFields(tags)).Warn(message)
	}
}

func Error(message, logType string, err error, tags map[string]string) {
	if Logger.Level >= logrus.ErrorLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = ErrorLevel.String()
		tags[Type] = logType
		message := fmt.Sprintf("%s - ERROR: %v", message, err)
		Logger.WithFields(formatFields(tags)).Error(message)
	}
}

func Panic(message, logType string, err error, tags map[string]string) {
	if Logger.Level >= logrus.PanicLevel {
		if tags == nil {
			tags = make(map[string]string)
		}
		tags[Level] = PanicLevel.String()
		tags[Type] = logType
		message := fmt.Sprintf("%s - PANIC: %v", message, err)
		Logger.WithFields(formatFields(tags)).Panic(message)
	}
}

func formatFields(tags map[string]string) logrus.Fields {
	fields := make(logrus.Fields)
	if tags != nil {
		for key, value := range tags {
			fields[strings.ReplaceAll(strings.TrimSpace(key), "_", "-")] = strings.ReplaceAll(strings.TrimSpace(value), "_", "-")
		}
	}
	return fields
}

func GetOut() io.Writer {
	return Logger.Out
}