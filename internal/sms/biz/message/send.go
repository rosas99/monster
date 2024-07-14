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
		templateM, err2 := b.ds.Templates().Get(ctx, rq.TemplateCode)
		if err2 != nil { // todo gorm record not found
			return nil, err2
		}
		jsonDataBytes, _ := json.Marshal(templateM)
		b.rds.Set(ctx, factory.WrapperTemplate(templateM.TemplateCode), jsonDataBytes, time.Hour*24)
		tm = templateM
	}

	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	_ = copier.Copy(&templateMsgRequest, rq)

	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)

	err = b.rule.CheckRules(ctx, cfgList)
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

	// todo log记录短信延时

	return &v1.CommonResponse{Code: tm.ID}, nil
	// todo 错误不为空，返回错误码
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
