package provider

import (
	"context"
	"fmt"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// DummyProvider is a struct represents a fake sms provider.
type DummyProvider struct{}

// NewDummyProvider returns a new provider for fake action.
func NewDummyProvider() *DummyProvider {
	return &DummyProvider{}
}

// Send do nothing
func (p *DummyProvider) Send(ctx context.Context, request *v1.TemplateMsgRequest) (TemplateMsgResponse, error) {
	// 模拟发送短信的逻辑，不实际发送
	fmt.Printf("Simulating message send via DummyProvider to %s\n", request.SendTime)
	return TemplateMsgResponse{}, nil
}
