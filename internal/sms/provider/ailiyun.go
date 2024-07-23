package provider

import (
	"context"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/writer"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
	ailiyunoptions "github.com/rosas99/monster/pkg/sdk/ailiyun"
)

// AILIYUNProvider is a struct represents a sms provider.
type AILIYUNProvider struct {
	rds               *redis.Client
	logger            *writer.Writer
	ailiyunSmsOptions *ailiyunoptions.SmsOptions
}

// NewAILIYUNProvider returns a new provider for aili cloud sms.
func NewAILIYUNProvider(rds *redis.Client, logger *writer.Writer, ailiyunSmsOptions *ailiyunoptions.SmsOptions) *AILIYUNProvider {
	return &AILIYUNProvider{
		rds:               rds,
		logger:            logger,
		ailiyunSmsOptions: ailiyunSmsOptions,
	}
}

// Send creates a sms client and sends sms by aili cloud.
func (p *AILIYUNProvider) Send(ctx context.Context, rq *v1.TemplateMsgRequest) (TemplateMsgResponse, error) {
	client, err := p.ailiyunSmsOptions.NewSmsClient()
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

	if tea.Int32Value(sendResp.StatusCode) != 200 {
		return TemplateMsgResponse{}, err
	}

	id := *sendResp.Body.BizId
	var history model.HistoryM
	history.MessageID = id
	p.logger.WriterHistory(&history)

	response := TemplateMsgResponse{
		Code:      *sendResp.Body.Code,
		Message:   *sendResp.Body.Message,
		BizId:     *sendResp.Body.BizId,
		RequestId: *sendResp.Body.RequestId,
	}
	return response, nil
}
