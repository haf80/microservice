package logger

import "log"

var L Logger = &DefaultLogger{}

func Warn(params ...interface{}) {
	L.Warn(params)
}

func Info(params ...interface{}) {
	L.Info(params)
}

func Error(params ...interface{}) {
	L.Error(params)
}

func Fatal(params ...interface{}) {
	L.Fatal(params)
}

type Logger interface {
	// Warn handles warning messages and logs
	Warn(...interface{})

	// Info handles informative logs and messages
	Info(...interface{})

	// Error handles messages and logs that point an error
	Error(...interface{})

	// Fatal logs errors and calls os.Exit()
	Fatal(...interface{})
}

// DefaultLogger is our default logger implementation
// TODO better style and more relevant syntax
type DefaultLogger struct {}

// NewDefaultLogger creates a new DefaultLogger
func NewDefaultLogger() *DefaultLogger {
	l := &DefaultLogger{}
	L = l
	return l
}

// Warn prints logs as a warning, each one in a separated line
func (l *DefaultLogger) Warn(params ...interface{}) {
	log.Println(params...)
}

// Info prints logs in info style, each one in a separated line
func (l *DefaultLogger) Info(params ...interface{}) {
	log.Println(params...)
}

// Error prints logs with an error style
func (l *DefaultLogger) Error(params ...interface{}) {
	log.Println(params...)
}

// Fatal prints logs and calls os.Exit()
func (l *DefaultLogger) Fatal(params ...interface{}) {
	log.Fatal(params...)
}
