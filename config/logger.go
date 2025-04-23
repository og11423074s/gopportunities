package config

import (
	"io"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	err     *log.Logger
	warning *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.LstdFlags)

	return &Logger{
		debug:   log.New(writer, colorBlue+"DEBUG: "+colorReset, logger.Flags()),
		info:    log.New(writer, colorGreen+"INFO: "+colorReset, logger.Flags()),
		err:     log.New(writer, colorRed+"ERROR: "+colorReset, logger.Flags()),
		warning: log.New(writer, colorYellow+"WARNING: "+colorReset, logger.Flags()),
		writer:  writer,
	}

}

// Create Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Create Formatted Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
