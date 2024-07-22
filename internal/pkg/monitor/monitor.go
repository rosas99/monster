package monitor

import (
	"github.com/rosas99/monster/pkg/log"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"github.com/segmentio/kafka-go"
	"sync"
)

var (
	once sync.Once
	cli  *impl
)

type impl struct {
	writer *kafka.Writer
}

// Monitor is a log.Monitor implementation that writes log messages to Kafka.
type Monitor struct {
	// enabled is an atomic boolean indicating whether the logger is enabled.
	enabled int32
	// writer is the Kafka writer used to write log messages.
	monitorWriter *kafka.Writer
}

// NewMonitor creates a new kafkaLogger instance.
func NewMonitor(monitorOpts *genericoptions.KafkaOptions) *impl {
	writer, err := monitorOpts.Writer()
	if err != nil {
		log.Fatalw("create kafka writer fail", "err", err)
	}

	once.Do(func() {
		cli = &impl{writer: writer}
	})
	return cli
}

// GetMonitor returns the globally initialized client.
func GetMonitor() *impl {
	return cli
}
