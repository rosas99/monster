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
	templateWriter *kafka.Writer
	uplinkWriter   *kafka.Writer
	uplinkWriter2  *kafka.Writer
	uplinkWriter3  *kafka.Writer
	ds             store.HistoryStore
}

// NewLogger creates a new kafkaLogger instance.
func NewLogger(templateOpts *genericoptions.KafkaOptions, uplinkOpts *genericoptions.KafkaOptions, uplinkOpts2 *genericoptions.KafkaOptions, uplinkOpts3 *genericoptions.KafkaOptions, ds store.HistoryStore) (*Logger, error) {
	templateWriter, err := templateOpts.Writer()
	if err != nil {
		return nil, err
	}
	uplinkWriter, err := uplinkOpts.Writer()
	if err != nil {
		return nil, err
	}
	uplinkWriter2, err := uplinkOpts2.Writer()
	if err != nil {
		return nil, err
	}
	uplinkWriter3, err := uplinkOpts3.Writer()
	if err != nil {
		return nil, err
	}

	logger := Logger{
		templateWriter: templateWriter,
		uplinkWriter:   uplinkWriter,
		uplinkWriter2:  uplinkWriter2,
		uplinkWriter3:  uplinkWriter3,
		ds:             ds,
	}
	return &logger, nil
}
