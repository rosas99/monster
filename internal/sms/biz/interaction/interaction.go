package interaction

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/writer"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// IBiz defines methods used to handle uplink message request.
type IBiz interface {
	AILIYUNUplink(ctx context.Context, rq *v1.AILIYUNUplinkListRequest) error
}

// interactionBiz struct implements the IBiz interface.
type interactionBiz struct {
	ds     store.IStore
	logger *writer.Writer
	rds    *redis.Client
	rule   *checker.RuleFactory
	idt    *idempotent.Idempotent
}

var _ IBiz = (*interactionBiz)(nil)

// New returns a new instance of interactionBiz.
func New(ds store.IStore, logger *writer.Writer, rds *redis.Client, idt *idempotent.Idempotent) *interactionBiz {
	return &interactionBiz{ds: ds, logger: logger, rds: rds, idt: idt}
}
