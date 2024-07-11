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

	writer *kafka.Writer
	ds     store.HistoryStore
}

// NewLogger creates a new kafkaLogger instance.
func NewLogger(kafkaOpts *genericoptions.KafkaOptions, ds store.HistoryStore) (*Logger, error) {
	writer, err := kafkaOpts.Writer()
	if err != nil {
		return nil, err
	}

	return &Logger{writer: writer, ds: ds}, nil
}
