//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"fmt"
	"github.com/op/go-logging"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var log = logging.MustGetLogger("lily")

func init() {
	log.ExtraCalldepth = 1
}

const (
	loggerPermissions = 0666
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
	defaultLoggerType   = "console"
	defaultLoggerPath   = ""
	defaultLoggerLayout = "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
	defaultLoggerLevel  = INFO
)

var (
	LOGGING_LEVELS = map[string]logging.Level{
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

type RotatoryWriter struct {
	doRotation bool
	file       *os.File
	filePath   string
}

func OpenRotatorFile(path string) io.Writer {
	out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, loggerPermissions)
	if err != nil {
		panic(
			fmt.Errorf(
				"Failed to open log file %s with permissions %d: %s", path, loggerPermissions, err,
			),
		)
	}
	rotator := &RotatoryWriter{doRotation: false, file: out, filePath: path}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func(rot *RotatoryWriter) {
		for {
			<-c
			rotator.doRotation = true
			Info("Rotate order caught for file '%s'", path)
		}
	}(rotator)
	return rotator
}

func (self *RotatoryWriter) Write(p []byte) (n int, err error) {
	if self.doRotation {
		self.doRotation = false
		self.file.Close()
		out, err := os.OpenFile(self.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, loggerPermissions)
		if err != nil {
			panic(
				fmt.Errorf(
					"Failed to open log file %s with permissions %d: %s", self.filePath, loggerPermissions, err,
				),
			)
		}
		self.file = out
	}
	return self.file.Write(p)
}
