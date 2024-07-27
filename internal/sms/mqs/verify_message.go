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

type VerifyMessageConsumer struct {
	ctx      context.Context
	idt      *idempotent.Idempotent
	logger   *writer.Writer
	provider *factory.ProviderFactory
}

func NewVerifyMessageConsumer(ctx context.Context, idt *idempotent.Idempotent, logger *writer.Writer, provider *factory.ProviderFactory) *CommonMessageConsumer {
	return &CommonMessageConsumer{
		ctx:      ctx,
		idt:      idt,
		logger:   logger,
		provider: provider,
	}
}

func (l *VerifyMessageConsumer) Consume(elem any) error {
	val := elem.(kafka.Message)

	var msg *types.TemplateMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		log.C(l.ctx).Errorf("Failed to unmarshal message: %v, value: %s", err, string(val.Value))
		return err
	}
	return l.handleSmsRequest(l.ctx, msg)

}

func (l *VerifyMessageConsumer) handleSmsRequest(ctx context.Context, msg *types.TemplateMsgRequest) error {

	if !l.idt.Check(ctx, msg.RequestId) {
		log.C(ctx).Errorf("Idempotent token check failed: %v", errors.New("idempotent token is invalid"))
		return errors.New("idempotent token is invalid")
	}
	log.C(ctx).Infof("Starting to process request: %v", msg.RequestId)

	historyM := model.HistoryM{}

	for _, provider := range msg.Providers {
		log.C(ctx).Infof("Processing provider: %v", provider)

		templateProvider, err := l.provider.GetSMSTemplateProvider(types.ProviderType(provider))
		if err != nil {
			log.C(ctx).Errorf("Failed to get SMS template provider: %v", err)
			continue
		}
		log.C(ctx).Infof("Sending message via provider: %v", provider)

		request := types.TemplateMsgRequest{}
		ret, err := templateProvider.Send(ctx, &request)
		log.C(ctx).Errorw(err, "send fail")

		if err != nil {
			log.C(ctx).Errorf("Failed to send SMS via provider %v: %v", provider, err)
			l.logger.WriterHistory(&historyM)
			continue
		}
		log.C(ctx).Infof("Message sent successfully via provider %v: bizId=%v", provider, ret.BizId)

		historyM.MessageID = ret.BizId
		l.logger.WriterHistory(&historyM)
		break
	}
	log.C(ctx).Infof("Finished processing request: %v", msg.RequestId)

	return nil
}
