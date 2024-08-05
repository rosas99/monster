package provider

import (
	"context"
	"fmt"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/internal/sms/writer"
	"github.com/rosas99/monster/pkg/log"
	ailiyunoptions "github.com/rosas99/monster/pkg/sdk/ailiyun"
)

// AILIYUNProvider is a struct represents a sms provider.
type AILIYUNProvider struct {
	typ    ProviderType
	rds    *redis.Client
	logger *writer.Writer
	client *dysmsapi.Client
}

// NewAILIYUNProvider returns a new provider for aili cloud sms.
func NewAILIYUNProvider(typ ProviderType, rds *redis.Client, logger *writer.Writer, ailiyunSmsOptions *ailiyunoptions.SmsOptions) *AILIYUNProvider {
	client, err := ailiyunSmsOptions.NewSmsClient()
	if err != nil {
		panic("unknown provider")
	}
	return &AILIYUNProvider{
		typ:    typ,
		rds:    rds,
		logger: logger,
		client: client,
	}
}

func (p *AILIYUNProvider) Type() ProviderType {
	return p.typ
}

// Send creates a sms client and sends sms by aili cloud.
func (p *AILIYUNProvider) Send(ctx context.Context, rq *types.TemplateMsgRequest) (TemplateMsgResponse, error) {
	log.C(ctx).Infof("Preparing to send SMS via AILIYUN for phone number: %v", rq.PhoneNumber)

	log.C(ctx).Infof("Created AILIYUN SMS client, preparing request")

	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(rq.PhoneNumber),
		SignName:      tea.String(rq.SignName),
		TemplateCode:  tea.String(rq.TemplateCode),
		TemplateParam: tea.String(rq.Content),
	}

	sendResp, err := p.client.SendSms(sendReq)

	if err != nil {
		log.C(ctx).Errorf("Failed to send SMS via AILIYUN: %v", err)
		return TemplateMsgResponse{}, err
	}
	log.C(ctx).Infof("Received response from AILIYUN, checking status code")

	if tea.Int32Value(sendResp.StatusCode) != 200 {
		log.C(ctx).Errorf("Non-200 status code received from AILIYUN: %v", sendResp.StatusCode)
		return TemplateMsgResponse{}, fmt.Errorf("non-200 status code: %v", sendResp.StatusCode)
	}

	id := *sendResp.Body.BizId
	var history model.HistoryM
	history.MessageID = id

	log.C(ctx).Infof("Recording history for message ID: %v", id)

	p.logger.WriterHistory(&history)

	log.C(ctx).Infof("SMS sent successfully via AILIYUN, preparing response")

	response := TemplateMsgResponse{
		Code:      *sendResp.Body.Code,
		Message:   *sendResp.Body.Message,
		BizId:     *sendResp.Body.BizId,
		RequestId: *sendResp.Body.RequestId,
	}
	log.C(ctx).Infof("Returning response from AILIYUN: %v", response)

	return response, nil
}
