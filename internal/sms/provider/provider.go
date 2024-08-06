package provider

import (
	"context"
	"fmt"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
)

type ProviderType string

const (
	ProviderTypeAliyun ProviderType = "aliyun"
	ProviderTypeDummy  ProviderType = "dummy"
)

// TemplateMsgResponse
type TemplateMsgResponse struct {
	Code      string
	Message   string
	BizId     string
	RequestId string
}

type Provider interface {
	// Type 返回 Provider 类型
	Type() ProviderType
	// Send 发送短信
	Send(ctx context.Context, rq *types.TemplateMsgRequest) (TemplateMsgResponse, error)
}

// ProviderFactory is a struct that acts as a factory for creating and managing instances
type ProviderFactory struct {
	providers map[types.ProviderType]Provider
}

// NewProviderFactory creates a new instance of ProviderFactory with an empty map of providers.
func NewProviderFactory() *ProviderFactory {
	return &ProviderFactory{
		providers: make(map[types.ProviderType]Provider),
	}
}

// RegisterProvider registers an SMS template provider.
func (f *ProviderFactory) RegisterProvider(providerType types.ProviderType, provider Provider) {
	f.providers[providerType] = provider
}

// GetSMSTemplateProvider retrieves an SMS template provider based on the given provider type.
func (f *ProviderFactory) GetSMSTemplateProvider(providerType types.ProviderType) (Provider, error) {
	log.Infof("Attempting to retrieve provider for type: %s", providerType)

	provider, exists := f.providers[providerType]
	if !exists {
		log.Errorf("No provider match for type: %s", providerType)
		return nil, fmt.Errorf("no provider match for type %s", providerType)
	}
	return provider, nil
}
