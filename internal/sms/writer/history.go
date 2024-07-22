package writer

import (
	"context"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/pkg/log"
)

// WriterHistory adds a new history record in the datastore.
func (l *Writer) WriterHistory(history *model.HistoryM) {
	err := l.historyStore.Create(context.Background(), history)
	if err != nil {
		log.Errorw(err, "Failed to create history messages")
	}
}
