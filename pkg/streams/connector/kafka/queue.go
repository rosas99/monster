package kafka

import (
	"context"
	"errors"
	"github.com/rosas99/monster/pkg/log"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"sync"
)

type ConsumeHandler interface {
	Consume(val any) error
}

type KQueue struct {
	consumer         *kafka.Reader
	handler          ConsumeHandler
	producerRoutines *sync.WaitGroup
	consumerRoutines *sync.WaitGroup
	forceCommit      bool
	channel          chan kafka.Message
	processors       int
	consumers        int
}

func NewKQueue(KafkaOptions *genericoptions.KafkaOptions, handler ConsumeHandler) (*KQueue, error) {
	r := kafka.ReaderConfig{
		Brokers:           KafkaOptions.Brokers,
		Topic:             KafkaOptions.Topic,
		GroupID:           KafkaOptions.ReaderOptions.GroupID,
		QueueCapacity:     KafkaOptions.ReaderOptions.QueueCapacity,
		MinBytes:          KafkaOptions.ReaderOptions.MinBytes,
		MaxBytes:          KafkaOptions.ReaderOptions.MaxBytes,
		MaxWait:           KafkaOptions.ReaderOptions.MaxWait,
		ReadBatchTimeout:  KafkaOptions.ReaderOptions.ReadBatchTimeout,
		HeartbeatInterval: KafkaOptions.ReaderOptions.HeartbeatInterval,
		CommitInterval:    KafkaOptions.ReaderOptions.CommitInterval,
		RebalanceTimeout:  KafkaOptions.ReaderOptions.RebalanceTimeout,
		StartOffset:       KafkaOptions.ReaderOptions.StartOffset,
		MaxAttempts:       KafkaOptions.ReaderOptions.MaxAttempts,
	}

	sink := &KQueue{
		consumer:         kafka.NewReader(r),
		handler:          handler,
		channel:          make(chan kafka.Message),
		producerRoutines: new(sync.WaitGroup),
		consumerRoutines: new(sync.WaitGroup),

		forceCommit: KafkaOptions.ForceCommit,
		processors:  KafkaOptions.Processors,
		consumers:   KafkaOptions.Consumers,
	}

	return sink, nil
}

func (c *KQueue) Start() {
	go c.startConsumers()
	go c.startProducers()

	c.producerRoutines.Wait()
	close(c.channel)
	c.consumerRoutines.Wait()
}

func (c *KQueue) Stop() {
	c.consumer.Close()
}
func (c *KQueue) startProducers() {
	for i := 0; i < c.consumers; i++ {
		c.producerRoutines.Add(1)
		go func() {
			defer c.producerRoutines.Done()
			for {
				msg, err := c.consumer.FetchMessage(context.Background())
				// io.EOF means consumer closed
				// io.ErrClosedPipe means committing messages on the consumer,
				// kafka will refire the messages on uncommitted messages, ignore
				if errors.Is(err, io.EOF) || errors.Is(err, io.ErrClosedPipe) {
					return
				}

				if err != nil {
					logx.Errorf("Error on reading message, %q", err.Error())
					continue
				}
				c.channel <- msg
			}
		}()
	}
}

func (c *KQueue) startConsumers() {
	for i := 0; i < c.processors; i++ {
		c.consumerRoutines.Add(1)
		go func() {
			defer c.consumerRoutines.Done()
			for msg := range c.channel {
				if err := c.handler.Consume(msg); err != nil {
					log.Errorf("consume: %s, error: %v", string(msg.Value), err)
					if !c.forceCommit {
						continue
					}
				}

				if err := c.consumer.CommitMessages(context.Background(), msg); err != nil {
					log.Errorf("commit failed, error: %v", err)
				}
			}
		}()

	}
}
