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
	"github.com/rosas99/monster/internal/sms/writer"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
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

	var msg *v1.UplinkMsgRequest
	err := json.Unmarshal(val.Value, &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}

	return l.handleSmsRequest(l.ctx, msg)
}

func (l *UplinkMessageConsumer) handleSmsRequest(ctx context.Context, msg *v1.UplinkMsgRequest) error {

	if !l.idt.Check(ctx, msg.RequestId) {
		return errors.New("idempotent token is invalid")
	}

	filter := make(map[string]any)
	filter["mobile"] = msg.PhoneNumber
	filter["content"] = msg.Content
	filter["receive_time"] = msg.SendTime
	count, _, _ := l.ds.Interactions().List(ctx, "", meta.WithFilter(filter))
	if count > 0 {
		// log record has existed
	}

	var interactionM model.InteractionM
	interactionM.InteractionID = uuid.New().String()
	interactionM.Mobile = msg.PhoneNumber
	interactionM.Content = msg.Content
	interactionM.Param = msg.DestCode
	interactionM.Provider = "AILIYUN"

	err := l.ds.Interactions().Create(ctx, &interactionM)
	if err != nil {
		// log
		return err
	}

	// todo 具体交互内容
	return nil
}
