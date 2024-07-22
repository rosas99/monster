package message

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/checker"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/writer"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// IBiz defines methods used to handle message-related functions.
type IBiz interface {
	Send(ctx context.Context, rq *v1.SendMessageRequest) error
	CodeVerify(ctx context.Context, rq *v1.VerifyCodeRequest) error
	AILIYUNReport(ctx context.Context, rq *v1.AILIYUNReportListRequest) error
}

// messageBiz struct implements the IBiz interface.
type messageBiz struct {
	ds     store.IStore
	logger *writer.Writer
	rds    *redis.Client
	rule   *checker.RuleFactory
	idt    *idempotent.Idempotent
}

var _ IBiz = (*messageBiz)(nil)

// New returns a new instance of messageBiz.
func New(ds store.IStore, logger *writer.Writer, rds *redis.Client, idt *idempotent.Idempotent) *messageBiz {
	return &messageBiz{ds: ds, logger: logger, rds: rds, idt: idt}
}
