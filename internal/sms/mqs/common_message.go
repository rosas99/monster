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
	"time"
)

type CommonMessageConsumer struct {
	ctx       context.Context
	idt       *idempotent.Idempotent
	logger    *writer.Writer
	providers map[string]factory.Provider
}

func NewCommonMessageConsumer(ctx context.Context, providers map[string]factory.Provider, idt *idempotent.Idempotent, logger *writer.Writer) *CommonMessageConsumer {

	return &CommonMessageConsumer{
		ctx:       ctx,
		idt:       idt,
		logger:    logger,
		providers: providers,
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

	historyM := model.HistoryM{
		Mobile:   msg.PhoneNumber, // Assuming PhoneNumber is the correct field name
		SendTime: time.Now(),      // Set current time as send time
		Content:  msg.Content,     // Assuming Content is the correct field name
		//MessageTemplateID: msg.TemplateCode,  // Assuming TemplateID is the correct field name
	}

	successful := false

	for _, provider := range msg.Providers {
		log.C(ctx).Infof("Attempting to use provider: %s", provider)
		providerIns, ok := l.providers[provider]
		if !ok {
			continue
		}
		ret, err := providerIns.Send(ctx, msg)

		if err != nil {
			log.C(ctx).Errorf("Failed to send SMS: %v", err)
			historyM.Status = "Failed"
			historyM.Message = err.Error() // Record the error message
		} else {
			log.C(ctx).Infof("SMS sent successfully: bizId=%v", ret.BizId)
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

	return nil
}
