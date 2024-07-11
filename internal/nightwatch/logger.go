package nightwatch

import (
	"github.com/rosas99/monster/pkg/log"
)

// cronLogger implement the cron.Logger interface.
type cronLogger struct{}

// newCronLogger returns a cron logger.
func newCronLogger() *cronLogger {
	return &cronLogger{}
}

// Info logs routine messages about cron's operation.
func (l *cronLogger) Info(msg string, keysAndValues ...any) {
	log.Infow(msg, keysAndValues...)
}

// Error logs an error condition.
func (l *cronLogger) Error(err error, msg string, keysAndValues ...any) {
	log.Errorw(err, msg, keysAndValues...)
}
