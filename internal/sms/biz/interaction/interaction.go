package interaction

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

type InteractionBiz interface {
	AILIYUNCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error)
}

// OrderBiz 接口的实现.
type interactionBiz struct {
	ds     store.IStore
	logger *logger.Logger
	rds    *redis.Client
	rule   *checker.RuleFactory
	idt    *idempotent.Idempotent
}

// 确保 orderBiz 实现了 OrderBiz 接口.
var _ InteractionBiz = (*interactionBiz)(nil)

// New 创建一个实现了 OrderBiz 接口的实例.
func New(ds store.IStore, logger *logger.Logger, rds *redis.Client, idt *idempotent.Idempotent) *interactionBiz {
	return &interactionBiz{ds: ds, logger: logger, rds: rds}
}

// todo 生成请求
// Create 是 OrderBiz 接口中 `Create` 方法的实现.
func (b *interactionBiz) AILIYUNCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error) {
	for index, i2 := range rq.AILIYUNCallbackList {
		fmt.Println(index, i2)

		// 直接在这里处理每条信息
		// todo 在原信息基础上补充接受到的时间和供应商类型
		// 构建interaction
		// 交互id使用自生成
	}
	// 根据手机号，内容，接收时间，查询数据库是否存在，存在则不保存，不存在则保存进数据库
	// 执行处理，如退订、兑换
	// 判断配置是否允许短信下行，允许的话回复用户并记录交互记录 //暂时不做
	return &v1.CommonResponse{Code: 0, Msg: "12"}, nil
}
