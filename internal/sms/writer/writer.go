package writer

import (
	"github.com/rosas99/monster/internal/sms/store"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"github.com/segmentio/kafka-go"
)

// Writer is a log.Writer implementation that writes log messages to Kafka.
type Writer struct {
	// enabled is an atomic boolean indicating whether the logger is enabled.
	enabled int32
	// writer is the Kafka writer used to write log messages.
	commonWriter *kafka.Writer
	verifyWriter *kafka.Writer
	uplinkWriter *kafka.Writer
	historyStore store.HistoryStore
}

// NewWriter creates a new kafkaLogger instance.
func NewWriter(commonOpts *genericoptions.KafkaOptions,
	verifyOpts *genericoptions.KafkaOptions,
	uplinkOpts *genericoptions.KafkaOptions,
	historyStore store.HistoryStore) (*Writer, error) {
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

	writer := Writer{
		commonWriter: commonWriter,
		verifyWriter: verifyWriter,
		uplinkWriter: uplinkWriter,
		historyStore: historyStore,
	}
	return &writer, nil
}
