package mqs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/provider"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/internal/sms/writer"
	"github.com/rosas99/monster/pkg/log"
	ailiyunoptions "github.com/rosas99/monster/pkg/sdk/ailiyun"
	"github.com/segmentio/kafka-go"
	"time"
)

type VerifyMessageConsumer struct {
	ctx               context.Context
	idt               *idempotent.Idempotent
	logger            *writer.Writer
	rds               *redis.Client
	ailiyunSmsOptions *ailiyunoptions.SmsOptions
}

func NewVerifyMessageConsumer(ctx context.Context, rds *redis.Client, idt *idempotent.Idempotent, logger *writer.Writer, ailiyunSmsOptions *ailiyunoptions.SmsOptions) *CommonMessageConsumer {
	return &CommonMessageConsumer{
		ctx:               ctx,
		idt:               idt,
		logger:            logger,
		rds:               rds,
		ailiyunSmsOptions: ailiyunSmsOptions,
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

	historyM := model.HistoryM{
		Mobile:   msg.PhoneNumber, // Assuming PhoneNumber is the correct field name
		SendTime: time.Now(),      // Set current time as send time
		Content:  msg.Content,     // Assuming Content is the correct field name
		//MessageTemplateID: msg.TemplateID,  // Assuming TemplateID is the correct field name
		Status: "Pending", // Initial status before sending
	}

	successful := false

	for _, provider := range msg.Providers {
		log.C(ctx).Infof("Processing provider: %v", provider)

		pins := factory.NewProvider(factory.ProviderTypeAliyun, l.rds, l.logger, l.ailiyunSmsOptions)

		ret, err := pins.Send(ctx, msg)

		if err != nil {
			log.C(ctx).Errorf("Failed to send SMS via provider %v: %v", provider, err)
			historyM.Status = "Failed"
			historyM.Message = err.Error()
		} else {
			log.C(ctx).Infof("Message sent successfully via provider %v: bizId=%v", provider, ret.BizId)
			historyM.Status = "Success"
			historyM.MessageID = ret.BizId
			historyM.Code = ret.Code
			historyM.Message = ret.Message
			successful = true
			break
		}
	}

	if successful {
		historyM.Status = "Success"
	} else {
		historyM.Status = "Failed"
	}

	l.logger.WriterHistory(&historyM)
	log.C(ctx).Infof("Finished processing request: %v", msg.RequestId)

	return nil
}
