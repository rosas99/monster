package provider

import (
	"context"
	"fmt"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
)

// TemplateMsgResponse
type TemplateMsgResponse struct {
	Code      string
	Message   string
	BizId     string
	RequestId string
}

// SMSTemplateProvider defines the SMS template sending interface.
type SMSTemplateProvider interface {
	Send(ctx context.Context, request *types.TemplateMsgRequest) (TemplateMsgResponse, error)
}

// ProviderFactory is a struct that acts as a factory for creating and managing instances
type ProviderFactory struct {
	providers map[types.ProviderType]SMSTemplateProvider
}

// NewProviderFactory creates a new instance of ProviderFactory with an empty map of providers.
func NewProviderFactory() *ProviderFactory {
	return &ProviderFactory{
		providers: make(map[types.ProviderType]SMSTemplateProvider),
	}
}

// RegisterProvider registers an SMS template provider.
func (f *ProviderFactory) RegisterProvider(providerType types.ProviderType, provider SMSTemplateProvider) {
	f.providers[providerType] = provider
}

// GetSMSTemplateProvider retrieves an SMS template provider based on the given provider type.
func (f *ProviderFactory) GetSMSTemplateProvider(providerType types.ProviderType) (SMSTemplateProvider, error) {
	log.Infof("Attempting to retrieve provider for type: %s", providerType)

	provider, exists := f.providers[providerType]
	if !exists {
		log.Errorf("No provider match for type: %s", providerType)
		return nil, fmt.Errorf("no provider match for type %s", providerType)
	}
	return provider, nil
}
