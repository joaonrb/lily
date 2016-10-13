package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

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
	CONSOLE = "console" // Console logging
	FILE    = "file"    // File Logging
)

const (
	CRITICAL = "critical" // Critical level
	ERROR    = "error"    // Error level
	WARNING  = "warning"  // Warning level
	INFO     = "info"     // Info level
	DEBUG    = "debug"    // Debug level
)

// Log default settings
const (
	defaultLoggerType   = "console"
	defaultLoggerPath   = ""
	defaultLoggerLayout = "%{level:.4s} %{time:2006-01-02 15:04:05.000} %{shortfile} %{message}"
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
)

// Critical logging
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
func LoadLogger(loggers []SLogger) {
	var out io.Writer
	goLoggers := make([]logging.Backend, 0)
	for _, loggerSettings := range loggers {
		switch loggerSettings.Type {
		case CONSOLE:
			out = os.Stdout
		case FILE:
			out = OpenRotatoryWriter(loggerSettings.Path)
		}
		logger := logging.NewLogBackend(out, "", 0)
		beFormatter := logging.NewBackendFormatter(logger, logging.MustStringFormatter(loggerSettings.Layout))
		beLevel := logging.AddModuleLevel(beFormatter)

		// Set Level
		lowerCaseLevel := strings.ToLower(loggerSettings.Level)
		levelNumber := LOGGING_LEVELS[lowerCaseLevel]
		beLevel.SetLevel(levelNumber, "")
		goLoggers = append(goLoggers, beLevel)
	}
	log.SetBackend(logging.MultiLogger(goLoggers...))
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
