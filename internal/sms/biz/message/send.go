package message

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/rosas99/monster/internal/sms/model"
	factory "github.com/rosas99/monster/internal/sms/store/redis"
	"github.com/rosas99/monster/internal/sms/types"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	"github.com/rosas99/monster/pkg/id"
	"github.com/rosas99/monster/pkg/log"
	"time"
)

// todo 生成请求
// Create 是 OrderBiz 接口中 `Create` 方法的实现.
func (b *messageBiz) Send(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	var templateMsgRequest types.TemplateMsgRequest
	templateMsgRequest.RequestId = b.idt.Token(ctx)
	// todo 先使用redis保存 后续再考虑使用本地缓存
	// todo 参考cache服务如何实现
	// 从本地缓存查询模板
	l := b.logger
	m := model.HistoryM{}

	templateM, err := b.ds.Templates().Get(ctx, rq.TemplateCode)
	if err != nil {
		l.LogHistory(&m)
	}

	// 如果是验证码，缓存验证码
	if templateM.Type == "VERIFICATION" {
		// todo 生成6位随机验证码
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
	err = b.rule.CheckRules(templateM, rq.Mobile, cfgList)
	// 这里只使用err
	if err != nil {
		l.LogHistory(&m)
	}

	// 根据类型发送短信到对应mq
	// 异步记录短信历史
	key := factory.WrapperCode(rq.TemplateCode, rq.Mobile)
	b.rds.Set(ctx, key, rq.Code, time.Hour*24)

	l.WriteCommonMessage(ctx, &templateMsgRequest)

	log.C(ctx).Infof("test")

	message := map[string]any{
		"test":  "value1",
		"other": 123,
	}
	l.LogKpi(message)
	// todo log记录短信延时
	return &v1.CreateTemplateResponse{OrderID: templateM.ID}, nil
}
