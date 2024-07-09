// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
	"k8s.io/klog/v2"

	"github.com/rosas99/monster/pkg/streams"
	"github.com/rosas99/monster/pkg/streams/flow"
)

type ConsumeHandler interface {
	Consume(val any) error
}

// KafkaSource represents an Apache Kafka source connector.
type KafkaSource struct {
	r           *kafka.Reader
	out         chan any
	ctx         context.Context
	cancelCtx   context.CancelFunc
	handler     ConsumeHandler
	parallelism uint
}

// NewKafkaSource returns a new KafkaSource instance.
func NewKafkaSource(ctx context.Context, config kafka.ReaderConfig, handler ConsumeHandler, parallelism uint) (*KafkaSource, error) {
	out := make(chan any)
	cctx, cancel := context.WithCancel(ctx)

	sink := &KafkaSource{
		r:           kafka.NewReader(config),
		out:         out,
		ctx:         cctx,
		cancelCtx:   cancel,
		handler:     handler,
		parallelism: parallelism,
	}

	go sink.init()
	return sink, nil
}

func (ks *KafkaSource) Start() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go ks.consume()

	select {
	case <-sigchan:
		ks.cancelCtx()
	case <-ks.ctx.Done():
	}

	close(ks.out)
	ks.r.Close()
}

// init starts the main loop.
func (ks *KafkaSource) init() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go ks.consume()

	select {
	case <-sigchan:
		ks.cancelCtx()
	case <-ks.ctx.Done():
	}

	close(ks.out)
	ks.r.Close()
}

func (ks *KafkaSource) consume() {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := ks.r.ReadMessage(ks.ctx)
		if err != nil {
			klog.ErrorS(err, "Failed to read message")
		}
		// 增加一个配置，可自行决定是否在业务逻辑处理失败时提交偏移量
		ks.out <- msg
		suc := <-ks.doStream()
		// 提交偏移量
		if suc {
			if err := ks.r.CommitMessages(context.Background(), msg); err != nil {
				// todo log.Fatal(err)
			}
		}
	}
}

// Via streams data through the given flow.
func (ks *KafkaSource) Via(_flow streams.Flow) streams.Flow {
	flow.DoStream(ks, _flow)
	return _flow
}

// Out returns an output channel for sending data.
func (ks *KafkaSource) Out() <-chan any {
	return ks.out
}

func (ks *KafkaSource) doStream() chan bool {
	suc := make(chan bool)
	sem := make(chan struct{}, ks.parallelism)
	for elem := range ks.out {
		sem <- struct{}{}
		go func(elem any) {
			defer func() { <-sem }()
			if err := ks.handler.Consume(elem); err != nil {
				// 处理错误
				// 提交消息等处理
				suc <- false
			} else {
				suc <- true
			}

			ks.out <- elem
		}(elem)
	}
	for i := 0; i < int(ks.parallelism); i++ {
		sem <- struct{}{}
	}
	close(ks.out)
	return suc
}

// KafkaSink represents an Apache Kafka sink connector.
type KafkaSink struct {
	ctx context.Context
	w   *kafka.Writer
	in  chan any
}

// NewKafkaSink returns a new KafkaSink instance.
func NewKafkaSink(ctx context.Context, config kafka.WriterConfig) (*KafkaSink, error) {
	sink := &KafkaSink{
		ctx: ctx,
		w:   kafka.NewWriter(config),
		in:  make(chan any),
	}

	go sink.init()
	return sink, nil
}

// init starts the main loop.
func (ks *KafkaSink) init() {
	for msg := range ks.in {
		var km kafka.Message
		switch m := msg.(type) {
		case []byte:
			km.Value = m
		case string:
			km.Value = []byte(m)
		case *kafka.Message:
			km = *m
		default:
			klog.V(1).InfoS("Unsupported message type", "message", m)
			continue
		}
		if err := ks.w.WriteMessages(ks.ctx, km); err != nil {
			klog.ErrorS(err, "Failed to write message")
		}
	}

	ks.w.Close()
}

// In returns an input channel for receiving data.
func (ks *KafkaSink) In() chan<- any {
	return ks.in
}
