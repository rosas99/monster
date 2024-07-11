package message

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
	"time"
)

// Send checks the template configuration and send the message to kafka queue.
func (b *messageBiz) Send(ctx context.Context, rq *v1.SendMessageRequest) (*v1.CommonResponse, error) {
	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)

	// todo 参考cache服务如何实现
	result, _ := b.rds.Get(ctx, factory.WrapperTemplateM(rq.TemplateCode)).Result()
	m := &model.TemplateM{}
	if result != "" {
		templateM := &model.TemplateM{}
		err := json.Unmarshal([]byte(result), &templateM)
		if err != nil {
			return nil, err
		}

		m = templateM
	} else {
		templateM, err := b.ds.Templates().Get(ctx, rq.TemplateCode)
		if err != nil {
			b.logger.LogHistory(&model.HistoryM{})
		}
		cache, err := json.Marshal(templateM)
		if err != nil {
			return nil, err
		}
		b.rds.Set(ctx, factory.WrapperTemplateM(templateM.TemplateCode), cache, time.Hour*24)
		m = templateM
	}

	if m.Type == types.VerificationMessage {
		rq.Code = id.RandomNumeric(6)
	}

	_ = copier.Copy(&templateMsgRequest, rq)

	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)
	if err != nil {
	}

	// todo 增加context
	err = b.rule.CheckRules(m, rq.Mobile, cfgList)
	if err != nil {
		historyM := model.HistoryM{}
		b.logger.LogHistory(&historyM)
	}

	key := factory.WrapperCode(rq.TemplateCode, rq.Mobile)
	b.rds.Set(ctx, key, rq.Code, time.Hour*24)

	b.logger.WriteCommonMessage(ctx, &templateMsgRequest)

	log.C(ctx).Infof("test")

	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}

	b.logger.LogKpi(message)

	// todo log记录短信延时

	return &v1.CommonResponse{Code: m.ID}, nil
}
