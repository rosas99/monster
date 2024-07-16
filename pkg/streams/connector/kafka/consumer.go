package kafka

import (
	"context"
	"errors"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"sync"
)

type ConsumeHandler interface {
	Consume(val any) error
}

type Consumer struct {
	r                *kafka.Reader
	handler          ConsumeHandler
	producerRoutines *sync.WaitGroup
	consumerRoutines *sync.WaitGroup
	forceCommit      bool
	channel          chan kafka.Message
	processors       int
	consumers        int
}

func NewConsumer(ctx context.Context, config kafka.ReaderConfig, handler ConsumeHandler, forceCommit bool) (*Consumer, error) {
	sink := &Consumer{
		r:           kafka.NewReader(config),
		handler:     handler,
		forceCommit: forceCommit,
	}

	return sink, nil
}

func (c *Consumer) Start() {
	go c.startConsumers()
	go c.startProducers()

	c.producerRoutines.Wait()
	close(c.channel)
	c.consumerRoutines.Wait()
}

func (c *Consumer) Stop() {
	c.r.Close()
}
func (c *Consumer) startProducers() {
	for i := 0; i < c.consumers; i++ {
		c.producerRoutines.Add(1)
		go func() {
			defer c.producerRoutines.Done()
			for {
				msg, err := c.r.FetchMessage(context.Background())
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
func (c *Consumer) startConsumers() {
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

				if err := c.r.CommitMessages(context.Background(), msg); err != nil {

				}
				log.Infof("")
			}
		}()

	}

}
