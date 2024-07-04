package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

func (l *Logger) LogMsg(msg *types.TemplateMsgRequest) {
	out, _ := json.Marshal(msg)
	fmt.Println(msg)
	if err := l.writer.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	} else {
		fmt.Println(string(out))
	}
}
