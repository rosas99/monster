package template

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/sms/store"
	v1 "github.com/rosas99/monster/pkg/api/sms/v1"
)

// IBiz defines methods used to handle template request.
type IBiz interface {
	Create(ctx context.Context, rq *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error)
	Get(ctx context.Context, id int64) (*v1.TemplateReply, error)
	List(ctx context.Context, rq *v1.ListTemplateRequest) (*v1.ListTemplateResponse, error)
	Update(ctx context.Context, id int64, rq *v1.UpdateTemplateRequest) error
	Delete(ctx context.Context, id int64) error
}

// templateBiz struct implements the IBiz interface.
type templateBiz struct {
	ds  store.IStore
	rds *redis.Client
}

func New(ds store.IStore, rds *redis.Client) *templateBiz {
	return &templateBiz{ds: ds, rds: rds}
}

// defines a group of constants for message configuration.
const (
	MessageCountForTemplatePerDay = "MESSAGE_COUNT_FOR_TEMPLATE_PER_DAY"
	MessageCountForMobilePerDay   = "MESSAGE_COUNT_FOR_MOBILE_PER_DAY"
	TimeIntervalForMobilePerDay   = "TIME_INTERVAL_FOR_MOBILE_PER_DAY"
)
