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
	commonWriter  *kafka.Writer
	verifyWriter  *kafka.Writer
	uplinkWriter  *kafka.Writer
	monitorWriter *kafka.Writer
	historyStore  store.HistoryStore
}

// NewLogger creates a new kafkaLogger instance.
func NewLogger(commonOpts *genericoptions.KafkaOptions,
	verifyOpts *genericoptions.KafkaOptions,
	uplinkOpts *genericoptions.KafkaOptions,
	monitorOpts *genericoptions.KafkaOptions,
	historyStore store.HistoryStore) (*Logger, error) {
	commonWriter, err := commonOpts.Writer()
	if err != nil {
		return nil, err
	}
	verifyWriter, err := verifyOpts.Writer()
	if err != nil {
		return nil, err
	}
	uplinkWriter, err := uplinkOpts.Writer()
	if err != nil {
		return nil, err
	}

	monitorWriter, err := monitorOpts.Writer()
	if err != nil {
		return nil, err
	}

	logger := Logger{
		commonWriter:  commonWriter,
		verifyWriter:  verifyWriter,
		uplinkWriter:  uplinkWriter,
		monitorWriter: monitorWriter,
		historyStore:  historyStore,
	}
	return &logger, nil
}
