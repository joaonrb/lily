//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"github.com/op/go-logging"
	"os"
	"io"
	"strings"
	"fmt"
	"os/signal"
	"syscall"
)

var log = logging.MustGetLogger("lily")

func init()  {
	log.ExtraCalldepth = 1
}

const (
	LOGGER_PERMISSIONS = 0666
)

const (
	CONSOLE = "console"
	FILE    = "file"
)

const (
	CRITICAL = "critical"
	ERROR    = "error"
	WARNING  = "warning"
	INFO     = "info"
	DEBUG    = "debug"
)

// Log default settings
const (
	DEFAULT_LOGGER_TYPE    = "console"
	DEFAULT_LOGGER_PATH    = ""
	DEFAULT_LOGGER_LAYOUT  = "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
	DEFAULT_LOGGER_LEVEL   = INFO
)

var (
	LOGGING_LEVELS = map[string]logging.Level {
		CRITICAL: logging.CRITICAL,
		ERROR:    logging.ERROR,
		WARNING:  logging.WARNING,
		INFO:     logging.INFO,
		DEBUG:    logging.DEBUG,
	}
)

func Critical(message string, args ...interface{}) {
	log.Criticalf(message, args...)
}

func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

func Warning(message string, args ...interface{}) {
	log.Warningf(message, args...)
}

func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Notice(message string, args ...interface{}) {
	log.Noticef(message, args...)
}

func Debug(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

func LoadLogger() {
	var out io.Writer
	loggers := make([]logging.Backend, 0)
	for _, loggerSettings := range Configuration.Loggers {
		switch loggerSettings.Type {
		case CONSOLE:
			out = os.Stdout
		case FILE:
			out = OpenRotatorFile(loggerSettings.Path)
		}
		logger := logging.NewLogBackend(out, "", 0)
		beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter(loggerSettings.Layout))
		beLevel := logging.AddModuleLevel(beFormatter)

		// Set Level
		lowerCaseLevel := strings.ToLower(loggerSettings.Level)
		levelNumber := LOGGING_LEVELS[lowerCaseLevel]
		beLevel.SetLevel(levelNumber, "")
		loggers = append(loggers, beLevel)
	}
	log.SetBackend(logging.MultiLogger(loggers...))
}

type RotatorWriter struct {
	doRotation bool
	file *os.File
	filePath string
}

func OpenRotatorFile(path string) io.Writer {
	out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, LOGGER_PERMISSIONS)
	if err != nil {
		panic(
			fmt.Errorf(
				"Failed to open log file %s with permissions %d: %s", path, LOGGER_PERMISSIONS, err,
			),
		)
	}
	rotator := &RotatorWriter{doRotation: false, file: out, filePath: path}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func (rot *RotatorWriter) {
		for {
			<- c
			rotator.doRotation = true
			Info("Rotate order caught for file '%s'", path)
		}
	}(rotator)
	return rotator
}

func (self *RotatorWriter) Write(p []byte) (n int, err error) {
	if self.doRotation {
		self.doRotation = false
		self.file.Close()
		out, err := os.OpenFile(self.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, LOGGER_PERMISSIONS)
		if err != nil {
			panic(
				fmt.Errorf(
					"Failed to open log file %s with permissions %d: %s", self.filePath, LOGGER_PERMISSIONS, err,
				),
			)
		}
		self.file = out
	}
	return self.file.Write(p)
}