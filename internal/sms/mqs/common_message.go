package consumer

import (
	"context"
	"encoding/json"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/model"
	providerFactory "github.com/rosas99/monster/internal/sms/provider"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageConsumer struct {
	ctx    context.Context
	idt    *idempotent.Idempotent
	logger *logger.Logger
}

func NewMessageConsumer(ctx context.Context, idt *idempotent.Idempotent, logger *logger.Logger) *MessageConsumer {
	return &MessageConsumer{
		ctx:    ctx,
		idt:    idt,
		logger: logger,
	}
}

func (l *MessageConsumer) Consume(elem any) error {
	val := elem.(kafka.Message)

	var msg *types.TemplateMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	return l.handleSmsRequest(l.ctx, msg)
}

func (l *MessageConsumer) handleSmsRequest(ctx context.Context, msg *types.TemplateMsgRequest) error {

	// 消息id
	ok := l.idt.Check(ctx, msg.RequestId)
	if !ok {
		// 消费失败
	}

	m := model.TemplateM{}
	providers := m.Providers
	for _, provider := range providers {
		providerFactory := providerFactory.NewProviderFactory()
		templateProvider, err := providerFactory.GetSMSTemplateProvider(types.ProviderType(provider))
		if err != nil {
			break
		}
		_, err = templateProvider.Send(types.TemplateMsgRequest{})
		if err != nil {
			continue
		}
		break
	}

	// todo 记录到history
	l.logger.LogHistory(&model.HistoryM{})
	return nil
}
