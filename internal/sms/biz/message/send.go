package message

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
	"time"

	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// Send checks the template configuration and send the message to kafka queue.
func (b *messageBiz) Send(ctx context.Context, rq *v1.SendMessageRequest) (*v1.CommonResponse, error) {
	templateCode := rq.TemplateCode
	tm := b.getTemplate(ctx, templateCode)
	if tm == nil {
		return nil, errors.New("")
	}

	cfgList := b.getCfgList(ctx, templateCode)
	if len(cfgList) == 0 {
		return nil, errors.New("")
	}

	err := b.rule.CheckRules(ctx, cfgList)
	if err != nil {
		b.log(rq, err, tm)
		// 记录错误码和错误类型
		return nil, err
	}

	if tm.Type == types.VerificationMessage {
		rq.Code = id.RandomNumeric(6)
		key := factory.WrapperCode(rq.TemplateCode, rq.Code)
		b.rds.Set(ctx, key, rq.Code, time.Hour*24)
	}

	// todo 分类型
	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	_ = copier.Copy(&templateMsgRequest, rq)
	err = b.logger.WriteMessage(ctx, &templateMsgRequest, tm.Type)
	if err != nil {
		log.C(ctx).Infof("test")
		b.log(rq, err, tm)
		return nil, err
	}

	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}

	b.logger.LogKpi(message)

	return &v1.CommonResponse{Code: tm.ID}, nil
}

func (b *messageBiz) getTemplate(ctx context.Context, templateCode string) *model.TemplateM {
	cache := b.fromTemplateCache(ctx, templateCode)

	if cache != nil {
		return cache
	}

	return b.fromTemplateDb(ctx, templateCode)

}

func (b *messageBiz) fromTemplateDb(ctx context.Context, templateCode string) *model.TemplateM {
	templateM, err := b.ds.Templates().GetByTemplateCode(ctx, templateCode)
	if err != nil { // todo gorm record not found
		return nil
	}

	marshal, _ := json.Marshal(templateM)
	b.rds.Set(ctx, factory.WrapperTemplate(templateM.TemplateCode), marshal, time.Hour*24)

	return templateM
}

func (b *messageBiz) fromTemplateCache(ctx context.Context, templateCode string) *model.TemplateM {
	tpCache, err := b.rds.Get(ctx, factory.WrapperTemplate(templateCode)).Result()
	if err != nil {
		return nil
	}

	tm := &model.TemplateM{}
	err = json.Unmarshal([]byte(tpCache), tm)
	if err != nil {
		return nil
	}

	return tm
}

func (b *messageBiz) getCfgList(ctx context.Context, templateCode string) []*model.ConfigurationM {
	var tm *model.TemplateM
	cache := b.fromCfgCache(ctx, templateCode)
	if len(cache) != 0 {
		return cache
	}

	return b.fromCfgDb(ctx, templateCode, tm)

}

func (b *messageBiz) fromCfgCache(ctx context.Context, templateCode string) []*model.ConfigurationM {
	cache, err := b.rds.Get(ctx, factory.WrapperTemplateCfg(templateCode)).Result()
	if err != nil || cache == "" {
		return nil
	}

	var cfgList []*model.ConfigurationM
	err = json.Unmarshal([]byte(cache), &cfgList)
	if err != nil {
		return nil
	}
	return cfgList
}

func (b *messageBiz) fromCfgDb(ctx context.Context, templateCode string, tm *model.TemplateM) []*model.ConfigurationM {
	_, list, err := b.ds.Configurations().List(ctx, templateCode)
	if err != nil { // todo gorm record not found
		return nil
	}

	marshal, _ := json.Marshal(list)
	b.rds.Set(ctx, factory.WrapperTemplateCfg(tm.TemplateCode), marshal, time.Hour*24)
	return list
}

func (b *messageBiz) log(rq *v1.SendMessageRequest, err error, m *model.TemplateM) {
	hm := model.HistoryM{
		Mobile:            maskPhone(rq.Mobile),
		SendTime:          time.Now(),
		Status:            types.ERROR_STATUS,
		Message:           err.Error(),
		Content:           m.Content,
		MessageTemplateID: m.ID,
	}
	b.logger.LogHistory(&hm)
}

func maskPhone(phone string) string {
	if len(phone) < 8 {
		return phone // 如果电话号码长度不足8位，则直接返回
	}
	mask := "****"
	return phone[:3] + mask + phone[7:]
}
