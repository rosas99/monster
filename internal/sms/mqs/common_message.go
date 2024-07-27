package mqs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/provider"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/internal/sms/writer"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

type CommonMessageConsumer struct {
	ctx      context.Context
	idt      *idempotent.Idempotent
	logger   *writer.Writer
	provider *factory.ProviderFactory
}

func NewCommonMessageConsumer(ctx context.Context, idt *idempotent.Idempotent, logger *writer.Writer, provider *factory.ProviderFactory) *CommonMessageConsumer {
	return &CommonMessageConsumer{
		ctx:      ctx,
		idt:      idt,
		logger:   logger,
		provider: provider,
	}
}

func (l *CommonMessageConsumer) Consume(elem any) error {
	val := elem.(kafka.Message)
	//log.C(l.ctx).Infof("Received message on topic %s with key: %s and partition: %d",
	//	val.TopicPartition.Topic, val.Key, val.TopicPartition.Partition)
	var msg *types.TemplateMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		log.C(l.ctx).Errorf("Failed to unmarshal message value: %v, error: %v", val.Value, err)
		return err
	}
	log.C(l.ctx).Infof("Successfully unmarshalled message: %v", msg)

	if err := l.handleSmsRequest(l.ctx, msg); err != nil {
		log.C(l.ctx).Errorf("Error handling SMS request: %v", err)
		return err
	}

	log.C(l.ctx).Infof("SMS request handled successfully")

	return nil

}

func (l *CommonMessageConsumer) handleSmsRequest(ctx context.Context, msg *types.TemplateMsgRequest) error {

	if !l.idt.Check(ctx, msg.RequestId) {
		log.C(ctx).Errorf("Idempotent token check failed: %v", errors.New("idempotent token is invalid"))
		return errors.New("idempotent token is invalid")
	}

	historyM := model.HistoryM{}

	for _, provider := range msg.Providers {
		log.C(ctx).Infof("Attempting to use provider: %s", provider)

		templateProvider, err := l.provider.GetSMSTemplateProvider(types.ProviderType(provider))
		if err != nil {
			log.C(ctx).Errorf("Failed to get SMS template provider: %v", err)
			continue
		}

		ret, err := templateProvider.Send(ctx, msg)

		if err != nil {
			log.C(ctx).Errorf("Failed to send SMS: %v", err)
			l.logger.WriterHistory(&historyM)
			continue
		}

		log.C(ctx).Infof("SMS sent successfully: bizId=%v", ret.BizId)

		historyM.MessageID = ret.BizId
		l.logger.WriterHistory(&historyM)
		break
	}

	return nil
}
