package interaction

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/logger"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// InteractionBiz defines methods used to handle uplink message request.
type InteractionBiz interface {
	AILIYUNCallback(ctx context.Context, rq *v1.AILIYUNCallbackListRequest) (*v1.CommonResponse, error)
}

// interactionBiz struct implements the InteractionBiz interface.
type interactionBiz struct {
	ds     store.IStore
	logger *logger.Logger
	rds    *redis.Client
	rule   *checker.RuleFactory
	idt    *idempotent.Idempotent
}

var _ InteractionBiz = (*interactionBiz)(nil)

// New returns a new instance of interactionBiz.
func New(ds store.IStore, logger *logger.Logger, rds *redis.Client, idt *idempotent.Idempotent) *interactionBiz {
	return &interactionBiz{ds: ds, logger: logger, rds: rds, idt: idt}
}
