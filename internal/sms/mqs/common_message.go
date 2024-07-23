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
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
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

	var msg *v1.TemplateMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		logx.Errorf("Consume val: %v error: %v", val, err)
		return err
	}
	return l.handleSmsRequest(l.ctx, msg)

}

func (l *CommonMessageConsumer) handleSmsRequest(ctx context.Context, msg *v1.TemplateMsgRequest) error {

	if !l.idt.Check(ctx, msg.RequestId) {
		return errors.New("idempotent token is invalid")
	}

	historyM := model.HistoryM{}

	for _, provider := range msg.Providers {
		templateProvider, err := l.provider.GetSMSTemplateProvider(types.ProviderType(provider))
		if err != nil {
			log.C(ctx).Errorw(err, "get fail")
			continue
		}

		ret, err := templateProvider.Send(ctx, msg)
		log.C(ctx).Errorw(err, "send fail")

		if err != nil {
			l.logger.WriterHistory(&historyM)
			continue
		}

		historyM.MessageID = ret.BizId
		l.logger.WriterHistory(&historyM)
		break
	}

	return nil
}
