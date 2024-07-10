package provider

import (
	"fmt"
	"github.com/rosas99/monster/internal/sms/types"
)

// TemplateMsgResponse 定义发送短信的响应结构
type TemplateMsgResponse struct {
	// todo 补充response
	MessageID string
	Error     error
}

// SMSTemplateProvider 接口定义
type SMSTemplateProvider interface {
	Send(request types.TemplateMsgRequest) (TemplateMsgResponse, error)
}

// ProviderFactory 管理短信服务提供者
type ProviderFactory struct {
	// 内部状态，比如映射表
	providers map[types.ProviderType]SMSTemplateProvider
}

// NewProviderFactory 创建并返回ProviderFactory的实例
func NewProviderFactory() *ProviderFactory {
	return &ProviderFactory{
		providers: make(map[types.ProviderType]SMSTemplateProvider),
	}
}

func (f *ProviderFactory) RegisterProvider(providerType types.ProviderType, provider SMSTemplateProvider) {
	f.providers[providerType] = provider
}

// 获取提供者示例
func (f *ProviderFactory) GetSMSTemplateProvider(providerType types.ProviderType) (SMSTemplateProvider, error) {
	provider, exists := f.providers[providerType]
	if !exists {
		return nil, fmt.Errorf("no provider match for type %d", providerType)
	}
	return provider, nil
}
