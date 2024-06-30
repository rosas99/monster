package provider

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/client"
	"github.com/rosas99/monster/internal/sms/types"
)

// WEProvider 结构体
type WEProvider struct {
	rds *redis.Client
}

// todo 依赖注入
func NewWEProvider(rds *redis.Client) *WEProvider {
	return &WEProvider{
		rds: rds,
	}
}

type AuthSuccess struct {
}

// Send 实现发送短信的方法
func (p *WEProvider) Send(rq types.TemplateMsgRequest) (TemplateMsgResponse, error) {
	// 这里应该是调用微信的API发送短信的逻辑
	fmt.Printf("Sending message via WEProvider to %s\n", rq.Request)
	// 返回示例响应

	url := "http://example.com/api"
	request := client.NewRequest(url)
	response, err := request.
		// 这里换成结构体
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		SetResult(&AuthSuccess{}).
		Post(url)
	if err != nil {

	}

	fmt.Print(response)

	if response.StatusCode() >= 400 {
		fmt.Printf("服务器返回错误状态码: %d\n", response.StatusCode())
		// 根据需要处理不同的错误状态码
	}

	return TemplateMsgResponse{MessageID: "123456"}, nil
}
