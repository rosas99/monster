package provider

import (
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/types"
)

// WEProvider 结构体
type AILIYUNProvider struct {
	rds    *redis.Client
	logger *logger.Logger
}

// todo 依赖注入
func NewAILIYUNProvider(rds *redis.Client, logger *logger.Logger) *AILIYUNProvider {
	return &AILIYUNProvider{
		rds:    rds,
		logger: logger,
	}
}
func (p *AILIYUNProvider) Send(rq types.TemplateMsgRequest) (TemplateMsgResponse, error) {

	return TemplateMsgResponse{MessageID: "123456"}, nil
}
