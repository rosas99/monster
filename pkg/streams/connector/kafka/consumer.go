package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"k8s.io/klog/v2"
)

type ConsumeHandler interface {
	Consume(val any) error
}

type Consumer struct {
	r           *kafka.Reader
	ctx         context.Context
	cancelCtx   context.CancelFunc
	handler     ConsumeHandler
	forceCommit bool
}

func NewConsumer(ctx context.Context, config kafka.ReaderConfig, handler ConsumeHandler, forceCommit bool) (*Consumer, error) {
	cctx, cancel := context.WithCancel(ctx)
	sink := &Consumer{
		r:           kafka.NewReader(config),
		ctx:         cctx,
		cancelCtx:   cancel,
		handler:     handler,
		forceCommit: forceCommit,
	}

	return sink, nil
}

func (ks *Consumer) Start() {
	go ks.consume()
}

func (ks *Consumer) Stop() {
	ks.cancelCtx()
	ks.r.Close()
}

func (ks *Consumer) consume() {
	for {
		msg, err := ks.r.ReadMessage(ks.ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to read message")
		}

		err = ks.handler.Consume(msg)
		if err != nil {
			// log error
			if !ks.forceCommit {
				continue
			}
		}

		err = ks.r.CommitMessages(context.Background(), msg)
		if err != nil {
			fmt.Println("commit messages fail")
		}
		fmt.Println("commit messages suc")

	}
}
