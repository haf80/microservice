package logger

import (
	"errors"
	"github.com/sirupsen/logrus"
)

// LogrusLogger is a logrus implementation of Logger interface
type LogrusLogger struct {}

// NewLogrusLogger returns new LogrusLogger
func NewLogrusLogger(lvl logrus.Level, f logrus.Formatter) *LogrusLogger {
	logrus.SetFormatter(f)
	logrus.SetLevel(lvl)
	lr := &LogrusLogger{}
	L = lr
	return lr
}

// Info calls logrus Info func
func (l *LogrusLogger) Info(params ...interface{}) {
	f, msg := mustExtractLogrusParams(params)
	logrus.WithFields(f).Info(msg)
}

// Warn calls logrus Warn func
func (l *LogrusLogger) Warn(params ...interface{}) {
	f, msg := mustExtractLogrusParams(params)
	logrus.WithFields(f).Warn(msg)
}

// Error calls logrus Error func
func (l *LogrusLogger) Error(params ...interface{}) {
	f, msg := mustExtractLogrusParams(params)
	logrus.WithFields(f).Error(msg)
}

// Fatal calls logrus Fatal func
func (l *LogrusLogger) Fatal(params ...interface{}) {
	f, msg := mustExtractLogrusParams(params)
	logrus.WithFields(f).Fatal(msg)
}

// mustExtractLogrusParams calls extractLogrusParams and ignores errors
// if there is any.
func mustExtractLogrusParams(params ...interface{}) (logrus.Fields, string) {
	f, s, _ := extractLogrusParams(params)
	return f, s
}

// extractLogrusParams tries to extract Fields and a message to pass them to
// logrus.WithFields. First param should be logrus.Fields. 2nd param is
// optional, and it supposed to be used as message. If 1 param is passed to
// this func, it returns an empty message.
func extractLogrusParams(params ...interface{}) (logrus.Fields, string, error) {
	p := params[0]
	f, err := convertToLogrusFields(p)
	if err != nil {
		return nil, "", err
	}

	if len(params) == 1 {
		return f, "", nil
	}

	if len(params) == 2 {
		if v, ok := params[1].(string); ok {
			return f, v, nil
		}
		return f, "", nil
	}

	return nil, "", errors.New("invalid params")
}

func convertToLogrusFields(param interface{}) (logrus.Fields, error) {
	switch t := param.(type) {
	case logrus.Fields:
		return t, nil
	case map[string]interface{}:
		var b map[string]interface{}
		// TODO stupid type cast xD. We should fix that.
		b = logrus.Fields(b)
		return b, nil
	}

	return nil, errors.New("invalid type")
}
