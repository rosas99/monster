package interaction

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/types"
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

// 放到队列比较合适，上行短信比较多可能处理不来
// Create 是 OrderBiz 接口中 `Create` 方法的实现.
func (b *interactionBiz) AILIYUNCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error) {
	for _, item := range rq.AILIYUNCallbackList {

		var msgRequest types.UplinkMsgRequest
		err := copier.Copy(msgRequest, item)
		if err != nil {
			return nil, err
		}
		b.logger.WriteUplinkMessage(ctx, &msgRequest)

	}
	// log kpi

	return &v1.CommonResponse{Code: 0, Msg: "success"}, nil
}
