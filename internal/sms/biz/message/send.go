package message

import (
	"context"
	"encoding/json"
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

	result, err := b.rds.Get(ctx, factory.WrapperTemplate(rq.TemplateCode)).Result()
	tm := &model.TemplateM{}
	if err != nil && result != "" {
		json.Unmarshal([]byte(result), tm)
	} else {
		templateM, err := b.ds.Templates().Get(ctx, rq.TemplateCode)
		if err != nil { // todo gorm record not found
			return nil, err
		}
		jsonDataBytes, _ := json.Marshal(templateM)
		b.rds.Set(ctx, factory.WrapperTemplate(templateM.TemplateCode), jsonDataBytes, time.Hour*24)
		tm = templateM
	}

	if tm.Type == types.VerificationMessage {
		rq.Code = id.RandomNumeric(6)
		// 缓存短信验证码
		key := factory.WrapperCode(rq.TemplateCode, rq.Code)
		b.rds.Set(ctx, key, rq.Code, time.Hour*24)
	}

	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	_ = copier.Copy(&templateMsgRequest, rq)

	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)

	err = b.rule.CheckRules(ctx, cfgList)
	if err != nil {
		b.log(rq)
		// 记录错误码和错误类型
		return nil, err
	}

	// todo 分类型
	err = b.logger.WriteCommonMessage(ctx, &templateMsgRequest)
	if err != nil {
		log.C(ctx).Infof("test")
		return nil, err
	}

	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}

	b.logger.LogKpi(message)
	// 组装好code和msg
	b.log(rq)

	// todo log记录短信延时

	return &v1.CommonResponse{Code: tm.ID}, nil
	// todo 错误不为空，返回错误码
}

func (b *messageBiz) log(rq *v1.SendMessageRequest) {
	hm := model.HistoryM{
		Mobile:   rq.Mobile,
		SendTime: time.Now(),
		Status:   "fail",
		// todo
		//Content:           templateM.Content,
		//MessageTemplateID: templateM.ID,
	}
	b.logger.LogHistory(&hm)
}
