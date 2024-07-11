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
	// todo 先使用redis保存 后续再考虑使用本地缓存
	// todo 参考cache服务如何实现
	// 从本地缓存查询模板
	result, err := b.rds.Get(ctx, factory.WrapperTemplateM(rq.TemplateCode)).Result()
	if err != nil {

	}
	m := model.TemplateM{}
	if result != "" {
		templateM := model.TemplateM{}
		err = json.Unmarshal([]byte(result), &templateM)
		if err != nil {

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
		m = *templateM
	}

	if m.Type == "VERIFICATION" {
		rq.Code = id.RandomNumeric(6)
	}

	// 组装短信
	_ = copier.Copy(&templateMsgRequest, rq)
	//templateMsgRequest.RequestId = strconv.FormatUint(snow.GenerateId(), 10)

	// 从本地缓存获取限流配置 // 忽略count返回
	_, cfgList, err := b.ds.Configurations().List(ctx, rq.TemplateCode)
	if err != nil {
	}

	// 规则校验
	err = b.rule.CheckRules(&m, rq.Mobile, cfgList)
	// 这里只使用err
	if err != nil {
		historyM := model.HistoryM{}
		b.logger.LogHistory(&historyM)
	}

	// 根据类型发送短信到对应mq
	// 异步记录短信历史
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
