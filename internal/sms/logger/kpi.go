package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

// LogKpi writes a log message for the api request.
func (l *Logger) LogKpi(messageMap map[string]any) {

	out, _ := json.Marshal(messageMap)

	fmt.Println(messageMap)
	if err := l.uplinkWriter.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	} else {
		fmt.Println(string(out))
	}
}
