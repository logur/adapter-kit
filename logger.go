// Package kit provides a Logur adapter for Go Kit logger.
package kit

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"logur.dev/adapter/kit/internal/keyvals"
)

// Logger is a Logur adapter for Go Kit logger.
type Logger struct {
	logger log.Logger
}

// New returns a new Logur logger.
// If logger is nil, a default instance is created.
func New(logger log.Logger) *Logger {
	if logger == nil {
		logger = log.NewNopLogger()
	}

	return &Logger{
		logger: logger,
	}
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	// Fall back to Debug
	l.Debug(msg, fields...)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	_ = level.Debug(l.logger).Log(append(l.keyvals(fields), "msg", msg)...)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	_ = level.Info(l.logger).Log(append(l.keyvals(fields), "msg", msg)...)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	_ = level.Warn(l.logger).Log(append(l.keyvals(fields), "msg", msg)...)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	_ = level.Error(l.logger).Log(append(l.keyvals(fields), "msg", msg)...)
}

func (l *Logger) keyvals(fields []map[string]interface{}) []interface{} {
	var kvs []interface{}
	if len(fields) > 0 {
		kvs = keyvals.FromMap(fields[0])
	}

	return kvs
}
