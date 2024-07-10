package kafka

import (
	"context"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
	"k8s.io/klog/v2"
)

type ConsumeHandler interface {
	Consume(val any) error
}

// KafkaSource represents an Apache Kafka source connector.
type Consumer struct {
	r           *kafka.Reader
	ctx         context.Context
	cancelCtx   context.CancelFunc
	handler     ConsumeHandler
	forceCommit bool
	stopCh      chan struct{}
}

// NewKafkaSource returns a new KafkaSource instance.
func NewConsumer(ctx context.Context, config kafka.ReaderConfig, handler ConsumeHandler, forceCommit bool) (*Consumer, error) {
	cctx, cancel := context.WithCancel(ctx)
	stopCh := make(chan struct{})
	sink := &Consumer{
		r:           kafka.NewReader(config),
		ctx:         cctx,
		cancelCtx:   cancel,
		handler:     handler,
		forceCommit: forceCommit,
		stopCh:      stopCh,
	}

	return sink, nil
}

func (ks *Consumer) Start() {
	go ks.consume(ks.stopCh)
}

func (ks *Consumer) Stop() {
	close(ks.stopCh)
	// 会取消关联的所有context的协程
	ks.cancelCtx()
	ks.r.Close()
}

func (ks *Consumer) consume(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh: // 监听停止信号
			log.Infof("Consumer is stopping gracefully...")
			return // 接收到停止信号，退出 consume 方法
		default:
			// the `ReadMessage` method blocks until we receive the next event
			msg, err := ks.r.ReadMessage(ks.ctx)
			if err != nil {
				klog.ErrorS(err, "Failed to read message")
			}

			if err := ks.handler.Consume(msg); err != nil {
				// log error
				if !ks.forceCommit {
					continue
				}
			}

			// 增加一个配置，可自行决定是否在业务逻辑处理失败时提交偏移量
			if err := ks.r.CommitMessages(context.Background(), msg); err != nil {
				// log commit error
			}

		}

	}
}
