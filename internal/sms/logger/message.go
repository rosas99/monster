package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

// WriteCommonMessage writes a log message for the common message.
func (l *Logger) WriteCommonMessage(ctx context.Context, msg *types.TemplateMsgRequest) {
	out, _ := json.Marshal(msg)
	fmt.Println(msg)
	if err := l.writer.WriteMessages(ctx, kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}

// WriteUplinkMessage writes a log message for the uplink message.
func (l *Logger) WriteUplinkMessage(ctx context.Context, msg *types.UplinkMsgRequest) {
	out, _ := json.Marshal(msg)
	fmt.Println(msg)
	if err := l.writer.WriteMessages(ctx, kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
