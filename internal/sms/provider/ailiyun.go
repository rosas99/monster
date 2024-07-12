package provider

import (
	"context"
	"fmt"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/sdk/ailiyun"
	"os"
)

// AILIYUNProvider is a struct represents a sms provider.
type AILIYUNProvider struct {
	rds    *redis.Client
	logger *logger.Logger
	url    string
}

// NewAILIYUNProvider returns a new provider for aili cloud sms.
func NewAILIYUNProvider(rds *redis.Client, logger *logger.Logger, url string) *AILIYUNProvider {
	return &AILIYUNProvider{
		rds:    rds,
		logger: logger,
		url:    url,
	}
}

// Send creates a sms client and sends sms by aili cloud.
func (p *AILIYUNProvider) Send(ctx context.Context, rq types.TemplateMsgRequest) (TemplateMsgResponse, error) {
	// 这里应该是调用微信的API发送短信的逻辑
	fmt.Printf("Sending message via WEProvider to %s\n", rq.SendTime)
	// 返回示例响应
	// todo 从配置获取
	fmt.Print("test config", p.url)
	client, err := ailiyun.CreateSmsClient(tea.String(os.Getenv("")), tea.String(os.Getenv("")))
	if err != nil {
		return TemplateMsgResponse{}, err
	}

	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(rq.PhoneNumber),
		SignName:      tea.String(rq.SignName),
		TemplateCode:  tea.String(rq.TemplateCode),
		TemplateParam: tea.String(rq.Content),
	}

	sendResp, err := client.SendSms(sendReq)
	if err != nil {
		return TemplateMsgResponse{}, err
	}
	id := *sendResp.Body.BizId
	fmt.Print(id)
	var history model.HistoryM
	history.MessageID = id
	p.logger.LogHistory(&history)

	// 组装code和msg
	// 根据err是否为nil组装状态码，存到history
	return TemplateMsgResponse{MessageID: "123456"}, nil
}
