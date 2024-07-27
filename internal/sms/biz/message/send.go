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

// Send is responsible for sending messages based on the request.
func (b *messageBiz) Send(ctx context.Context, rq *v1.SendMessageRequest) error {
	log.C(ctx).Infof("Starting message send process for request: %v", rq)

	tm := b.getTemplate(ctx, rq.TemplateCode)
	if tm == nil {
		log.C(ctx).Errorf("Template not found for template code: %s", rq.TemplateCode)
		return errors.New("template not found")
	}
	log.C(ctx).Infof("Template found: %v", tm)

	cfgList := b.getCfgList(ctx, rq.TemplateCode)
	if len(cfgList) == 0 {
		log.C(ctx).Errorf("No configurations found for template code: %s", rq.TemplateCode)
		return errors.New("no configurations found")
	}
	log.C(ctx).Infof("Configurations found: %v", cfgList)
	err := b.rule.CheckRules(ctx, cfgList)
	if err != nil {
		log.C(ctx).Errorf("Rule check failed: %v", err)
		b.log(rq, err, tm)
		return err
	}
	log.C(ctx).Infof("Rules checked successfully")
	if tm.Type == types.VerificationMessage {
		rq.Code = id.RandomNumeric(6)
		key := wrapper.WrapperCode(rq.TemplateCode, rq.Code)
		b.rds.Set(ctx, key, rq.Code, time.Hour*24)
		log.C(ctx).Infof("Verification code set: %s", key)
	}

	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	_ = copier.Copy(&templateMsgRequest, rq)
	err = b.logger.WriteMessage(ctx, &templateMsgRequest, tm.Type)
	if err != nil {
		log.C(ctx).Errorf("Failed to write message to logger: %v", err)
		b.log(rq, err, tm)
		b.log(rq, err, tm)
		return err
	}
	log.C(ctx).Infof("Message sent successfully")
	return nil
}

func (b *messageBiz) getTemplate(ctx context.Context, templateCode string) *model.TemplateM {
	log.C(ctx).Infof("Fetching template for templateCode: %s", templateCode)

	cache, _ := b.rds.Get(ctx, wrapper.WrapperTemplate(templateCode)).Result()
	if cache != "" {
		tm := &model.TemplateM{}
		if err := json.Unmarshal([]byte(cache), tm); err != nil {
			log.C(ctx).Errorf("Error unmarshalling cache data for templateCode: %s, error: %v", templateCode, err)
			return nil
		}

		log.C(ctx).Infof("Template fetched from cache for templateCode: %s", templateCode)
		return tm
	}

	tm, _ := b.ds.Templates().GetByTemplateCode(ctx, templateCode)
	if tm != nil {
		marshal, _ := json.Marshal(tm)
		b.rds.Set(ctx, wrapper.WrapperTemplate(tm.TemplateCode), marshal, time.Hour*24)
		log.C(ctx).Infof("Template fetched from database and cached for templateCode: %s", templateCode)
		return tm
	}

	log.C(ctx).Warnf("Template not found for templateCode: %s", templateCode)
	return nil
}

func (b *messageBiz) getCfgList(ctx context.Context, templateCode string) []*model.ConfigurationM {
	log.C(ctx).Infof("Fetching configurations for template code: %s", templateCode)

	cache, _ := b.rds.Get(ctx, wrapper.WrapperTemplateCfg(templateCode)).Result()
	if cache != "" {
		var cfgList []*model.ConfigurationM
		if err := json.Unmarshal([]byte(cache), &cfgList); err != nil {
			log.C(ctx).Errorf("Error unmarshalling cache data for configurations of template code: %s, error: %v", templateCode, err)
			return nil
		}

		log.C(ctx).Infof("Configurations fetched from cache for template code: %s", templateCode)
		return cfgList
	}

	_, list, _ := b.ds.Configurations().List(ctx, templateCode)
	if len(list) <= 0 {
		log.C(ctx).Warnf("No configurations found for template code: %s", templateCode)
		return nil
	}

	marshal, _ := json.Marshal(list)
	b.rds.Set(ctx, wrapper.WrapperTemplateCfg(templateCode), marshal, time.Hour*24)
	log.C(ctx).Infof("Configurations fetched from database and cached for template code: %s", templateCode)
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
