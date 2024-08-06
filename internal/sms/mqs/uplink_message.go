package mqs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/internal/sms/writer"
	"github.com/rosas99/monster/pkg/log"
	"github.com/segmentio/kafka-go"
)

type UplinkMessageConsumer struct {
	ctx    context.Context
	idt    *idempotent.Idempotent
	logger *writer.Writer
	ds     store.IStore
}

func NewUplinkMessageConsumer(ctx context.Context, ds store.IStore, idt *idempotent.Idempotent, logger *writer.Writer) *UplinkMessageConsumer {
	return &UplinkMessageConsumer{
		ctx:    ctx,
		idt:    idt,
		logger: logger,
		ds:     ds,
	}
}

func (l *UplinkMessageConsumer) Consume(elem any) error {
	val := elem.(kafka.Message)

	var msg *types.UplinkMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		log.C(l.ctx).Errorf("Failed to unmarshal message: %v, value: %s", err, string(val.Value))
		return err
	}

	log.C(l.ctx).Infof("Uplink message consumed: %v", msg)
	return l.handleSmsRequest(l.ctx, msg)
}

func (l *UplinkMessageConsumer) handleSmsRequest(ctx context.Context, msg *types.UplinkMsgRequest) error {

	if !l.idt.Check(ctx, msg.RequestId) {
		log.C(ctx).Errorf("Idempotent token check failed: %v", errors.New("idempotent token is invalid"))
		return errors.New("idempotent token is invalid")
	}
	log.C(ctx).Infof("Checking for existing interaction records for mobile: %v", msg.PhoneNumber)

	filter := make(map[string]any)
	filter["mobile"] = msg.PhoneNumber
	filter["content"] = msg.Content
	filter["receive_time"] = msg.SendTime
	count, _, _ := l.ds.Interactions().List(ctx, meta.WithFilter(filter))
	if count > 0 {
		log.C(ctx).Infof("Interaction record already exists for mobile: %v", msg.PhoneNumber)
	}

	var interactionM model.InteractionM
	interactionM.InteractionID = uuid.New().String()
	interactionM.Mobile = msg.PhoneNumber
	interactionM.Content = msg.Content
	interactionM.Param = msg.DestCode
	interactionM.Provider = "AILIYUN"
	log.C(ctx).Infof("Creating new interaction record: %v", interactionM)

	err := l.ds.Interactions().Create(ctx, &interactionM)
	if err != nil {
		log.C(ctx).Errorf("Failed to create interaction record: %v", err)
		return err
	}
	log.C(ctx).Infof("Interaction record created successfully: %v", interactionM.InteractionID)

	// todo 具体交互内容
	return nil
}
