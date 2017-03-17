// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"errors"
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
	CONSOLE = "console" // console logging
	FILE    = "file"    // file Logging
)

const (
	CRITICAL = "critical" // critical level
	ERROR    = "error"    // error level
	WARNING  = "warning"  // warning level
	INFO     = "info"     // info level
	DEBUG    = "debug"    // debug level
)

// Log default settings
const (
	defaultLoggerType   = "console"
	defaultLoggerPath   = ""
	defaultLoggerFormat = "%{color}%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile}%{color:reset} %{message}"
	defaultLoggerLevel  = INFO
)

var (
	// Logging map of levels
	LOGGING_LEVELS = map[string]logging.Level{
		CRITICAL: logging.CRITICAL,
		ERROR:    logging.ERROR,
		WARNING:  logging.WARNING,
		INFO:     logging.INFO,
		DEBUG:    logging.DEBUG,
	}
	defaultLogger = []interface{}{
		map[string]interface{}{
			"type":   defaultLoggerType,
			"format": defaultLoggerFormat,
			"level":  defaultLoggerLevel,
		},
	}
)

// Error logging
func Critical(message string, args ...interface{}) {
	log.Criticalf(message, args...)
}

// Error logging
func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

// Warning logging
func Warning(message string, args ...interface{}) {
	log.Warningf(message, args...)
}

// Info logging
func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

// Notice logging
func Notice(message string, args ...interface{}) {
	log.Noticef(message, args...)
}

// Debug logging
func Debug(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

// Load logger
func LoadLogger() error {
	var out io.Writer
	var loggers []interface{}

	value, ok := Settings["loggers"]
	if !ok {
		loggers = defaultLogger
	} else if loggers, ok = value.([]interface{}); !ok {
		loggers = defaultLogger
	}

	goLoggers := []logging.Backend{}
	for _, loggerInterface := range loggers {
		loggerSettings := loggerInterface.(map[interface{}]interface{})
		switch loggerSettings["type"].(string) {
		case CONSOLE:
			out = os.Stdout
		case FILE:
			if path, ok := loggerSettings["path"].(string); ok {
				out = OpenRotatoryWriter(path)
			} else {
				return errors.New("Path is not defined for type file.")
			}
		default:
			return errors.New("Type is not defined or is wrong.")
		}
		logger := logging.NewLogBackend(out, "", 0)
		format, ok := loggerSettings["format"].(string)
		if !ok {
			format = defaultLoggerFormat
		}
		beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter(format))
		beLevel := logging.AddModuleLevel(beFormatter)

		// Set Level
		level, ok := loggerSettings["level"].(string)
		if !ok {
			level = defaultLoggerLevel
		}
		lowerCaseLevel := strings.ToLower(level)
		levelNumber := LOGGING_LEVELS[lowerCaseLevel]
		beLevel.SetLevel(levelNumber, "")
		goLoggers = append(goLoggers, beLevel)
	}
	log.SetBackend(logging.MultiLogger(goLoggers...))
	return nil
}

// Rotate writer
type RotatoryWriter struct {
	doRotation bool
	file       *os.File
	filePath   string
}

// Open the writer
func OpenRotatoryWriter(path string) io.Writer {
	out, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, loggerPermissions)
	if err != nil {
		panic(
			fmt.Errorf(
				"Failed to open log file %s with permissions %d: %s", path, loggerPermissions, err,
			),
		)
	}
	rotatoryWriter := &RotatoryWriter{doRotation: false, file: out, filePath: path}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func(rot *RotatoryWriter) {
		for {
			<-c
			rotatoryWriter.doRotation = true
			Info("Rotate order caught for file '%s'", path)
		}
	}(rotatoryWriter)
	return rotatoryWriter
}

// Write on file
func (writer *RotatoryWriter) Write(p []byte) (n int, err error) {
	if writer.doRotation {
		writer.doRotation = false
		writer.file.Close()
		out, err := os.OpenFile(writer.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, loggerPermissions)
		if err != nil {
			panic(
				fmt.Errorf(
					"Failed to open log file %s with permissions %d: %s", writer.filePath, loggerPermissions, err,
				),
			)
		}
		writer.file = out
	}
	return writer.file.Write(p)
}
