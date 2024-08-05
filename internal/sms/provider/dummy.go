package provider

import (
	"context"
	"github.com/rosas99/monster/internal/sms/types"
	"github.com/rosas99/monster/pkg/log"
)

// DummyProvider is a struct represents a fake sms provider.
type DummyProvider struct {
	typ ProviderType
}

// NewDummyProvider returns a new provider for fake action.
func NewDummyProvider(typ ProviderType) *DummyProvider {
	return &DummyProvider{
		typ: typ,
	}
}

func (p *DummyProvider) Type() ProviderType {
	return p.typ
}

// Send do nothing
func (p *DummyProvider) Send(ctx context.Context, request *types.TemplateMsgRequest) (TemplateMsgResponse, error) {
	log.C(ctx).Infof("Simulating message send via DummyProvider to %s", request.PhoneNumber)

	// Since this is a dummy provider, no real action is taken here.
	// The response is returned as if the operation was successful.
	return TemplateMsgResponse{}, nil
}
