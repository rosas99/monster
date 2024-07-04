package logger

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/pkg/log"
)

// LogModel writes a log message for the policy model.
func (l *Logger) LogHistory(history *model.HistoryM) {
	err := l.ds.Create(context.Background(), history)
	if err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
