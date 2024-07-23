package message

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	wrapper "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
	"time"

	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Send checks the template configuration and send the message to kafka queue.
func (b *messageBiz) Send(ctx context.Context, rq *v1.SendMessageRequest) error {
	tm := b.getTemplate(ctx, rq.TemplateCode)
	if tm == nil {
		return errors.New("")
	}

	cfgList := b.getCfgList(ctx, rq.TemplateCode)
	if len(cfgList) == 0 {
		return errors.New("")
	}

	err := b.rule.CheckRules(ctx, cfgList)
	if err != nil {
		b.log(rq, err, tm)
		return err
	}

	if tm.Type == types.VerificationMessage {
		rq.Code = id.RandomNumeric(6)
		key := wrapper.WrapperCode(rq.TemplateCode, rq.Code)
		b.rds.Set(ctx, key, rq.Code, time.Hour*24)
	}

	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	_ = copier.Copy(&templateMsgRequest, rq)
	err = b.logger.WriteMessage(ctx, &templateMsgRequest, tm.Type)
	if err != nil {
		log.C(ctx).Infof("test")
		b.log(rq, err, tm)
		return err
	}

	return nil
}

func (b *messageBiz) getTemplate(ctx context.Context, templateCode string) *model.TemplateM {
	tpCache, _ := b.rds.Get(ctx, wrapper.WrapperTemplate(templateCode)).Result()
	if tpCache != "" {
		tm := &model.TemplateM{}
		if err := json.Unmarshal([]byte(tpCache), tm); err != nil {
			return nil
		}

		return tm
	}

	tm, _ := b.ds.Templates().GetByTemplateCode(ctx, templateCode)
	if tm != nil {
		marshal, _ := json.Marshal(tm)
		b.rds.Set(ctx, wrapper.WrapperTemplate(tm.TemplateCode), marshal, time.Hour*24)
		return tm
	}
	return nil
}

func (b *messageBiz) getCfgList(ctx context.Context, templateCode string) []*model.ConfigurationM {
	cache, _ := b.rds.Get(ctx, wrapper.WrapperTemplateCfg(templateCode)).Result()
	if cache != "" {
		var cfgList []*model.ConfigurationM
		if err := json.Unmarshal([]byte(cache), &cfgList); err != nil {
			return nil
		}

		return cfgList
	}

	_, list, _ := b.ds.Configurations().List(ctx, templateCode)
	if len(list) <= 0 {
		return nil
	}

	marshal, _ := json.Marshal(list)
	b.rds.Set(ctx, wrapper.WrapperTemplateCfg(templateCode), marshal, time.Hour*24)
	return list
}

func (b *messageBiz) log(rq *v1.SendMessageRequest, err error, m *model.TemplateM) {
	hm := model.HistoryM{
		Mobile:            maskPhone(rq.Mobile),
		SendTime:          time.Now(),
		Status:            types.ErrorStatus,
		Message:           err.Error(),
		Content:           m.Content,
		MessageTemplateID: m.ID,
	}
	b.logger.WriterHistory(&hm)
}

func maskPhone(phone string) string {
	if len(phone) < 8 {
		return phone
	}
	mask := "****"
	return phone[:3] + mask + phone[7:]
}
