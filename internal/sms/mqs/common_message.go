package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/provider"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommonMessageConsumer struct {
	ctx      context.Context
	idt      *idempotent.Idempotent
	logger   *logger.Logger
	provider *factory.ProviderFactory
}

func NewCommonMessageConsumer(ctx context.Context, idt *idempotent.Idempotent, logger *logger.Logger, provider *factory.ProviderFactory) *CommonMessageConsumer {
	return &CommonMessageConsumer{
		ctx:      ctx,
		idt:      idt,
		logger:   logger,
		provider: provider,
	}
}

func (l *CommonMessageConsumer) Consume(elem any) error {
	val := elem.(kafka.Message)

	var msg *types.TemplateMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	return l.handleSmsRequest(l.ctx, msg)

}

func (l *CommonMessageConsumer) handleSmsRequest(ctx context.Context, msg *types.TemplateMsgRequest) error {

	// 消息id
	ok := l.idt.Check(ctx, msg.RequestId)
	if !ok {
		return errors.New("message repeat")
	}

	m := model.TemplateM{}
	providers := m.Providers
	for _, provider := range providers {
		templateProvider, err := l.provider.GetSMSTemplateProvider(types.ProviderType(provider))
		if err != nil {
			break
		}
		ret, err := templateProvider.Send(ctx, types.TemplateMsgRequest{})
		if err != nil {
			continue
		}
		historyM := model.HistoryM{
			MessageID: ret.MessageID,
		}
		// todo 记录到history
		// 从响应获取bizid，关联到history
		l.logger.LogHistory(&historyM)
		break
	}

	return nil
}
