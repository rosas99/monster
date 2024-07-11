package logger

import (
	"github.com/rosas99/monster/internal/sms/store"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"github.com/segmentio/kafka-go"
)

// Logger is a log.Logger implementation that writes log messages to Kafka.
type Logger struct {
	// enabled is an atomic boolean indicating whether the logger is enabled.
	enabled int32
	// writer is the Kafka writer used to write log messages.

	// 不同的队列使用不同的writer
	writer  *kafka.Writer
	writer2 *kafka.Writer
	ds      store.HistoryStore
}

// NewLogger creates a new kafkaLogger instance.
func NewLogger(kafkaOpts *genericoptions.KafkaOptions, kafkaOpts2 *genericoptions.KafkaOptions, ds store.HistoryStore) (*Logger, error) {
	writer, err := kafkaOpts.Writer()
	if err != nil {
		return nil, err
	}
	writer2, err := kafkaOpts2.Writer()
	if err != nil {
		return nil, err
	}

	return &Logger{writer: writer, writer2: writer2, ds: ds}, nil
}
