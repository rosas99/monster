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
func (l *Logger) WriteMessage(ctx context.Context, msg *types.TemplateMsgRequest, messageType string) error {
	out, _ := json.Marshal(msg)
	if messageType == types.VerificationMessage {
		return l.writer2.WriteMessages(ctx, kafka.Message{Value: out})
	} else {
		return l.writer.WriteMessages(ctx, kafka.Message{Value: out})
	}

}

func (l *Logger) WriteVerifyMessage(ctx context.Context, msg *types.TemplateMsgRequest) error {
	out, _ := json.Marshal(msg)
	fmt.Println(msg)
	if err := l.writer2.WriteMessages(ctx, kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
		return err
	}
	return nil
}

// WriteUplinkMessage writes a log message for the uplink message.
func (l *Logger) WriteUplinkMessage(ctx context.Context, msg *types.UplinkMsgRequest) {
	out, _ := json.Marshal(msg)
	fmt.Println(msg)
	if err := l.writer2.WriteMessages(ctx, kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
