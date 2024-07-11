package logger

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/pkg/log"
)

// LogHistory adds a new secret record in the datastore.
func (l *Logger) LogHistory(history *model.HistoryM) {
	err := l.ds.Create(context.Background(), history)
	if err != nil {
		log.Errorw(err, "Failed to create history messages")
	}
}
