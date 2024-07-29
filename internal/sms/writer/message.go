package writer

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

const (
	VerificationMessage = "VERIFICATION_MESSAGE"
)

// WriteMessage writes a log message for the common message.
func (l *Writer) WriteMessage(ctx context.Context, msg *types.TemplateMsgRequest, messageType string) error {
	out, _ := json.Marshal(msg)
	if messageType == VerificationMessage {
		return l.commonWriter.WriteMessages(ctx, kafka.Message{Value: out})
	} else {
		return l.verifyWriter.WriteMessages(ctx, kafka.Message{Value: out})
	}

}

// WriteUplinkMessage writes a log message for the uplink message.
func (l *Writer) WriteUplinkMessage(ctx context.Context, msg *types.UplinkMsgRequest) {
	out, _ := json.Marshal(msg)
	if err := l.uplinkWriter.WriteMessages(ctx, kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}
