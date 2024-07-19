package logger

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

type AuditMessage struct {
}

func (l *Logger) LogAuditMessage(data AuditMessage) {
	out, _ := json.Marshal(data)
	if err := l.monitorWriter.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}

// LogKpi writes a log message for the api request.
func (l *Logger) LogKpi(messageMap map[string]any) {
	out, _ := json.Marshal(messageMap)
	if err := l.monitorWriter.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
