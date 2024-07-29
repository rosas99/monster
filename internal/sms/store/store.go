package store

//go:generate mockgen -destination mock_store.go -package store github.com/rosas99/monster/internal/fakeserver/store IStore,OrderStore

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
)

// IStore is an interface that represents methods
// required to be implemented by a Store implementation.
type IStore interface {
	TX(context.Context, func(ctx context.Context) error) error
	Templates() TemplateStore
	Configurations() ConfigurationStore
	Histories() HistoryStore
	Interactions() InteractionStore
}

// TemplateStore defines the interface for managing template data storage.
type TemplateStore interface {
	Create(ctx context.Context, order *model.TemplateM) error
	Get(ctx context.Context, id int64) (*model.TemplateM, error)
	GetByTemplateCode(ctx context.Context, templateCode string) (*model.TemplateM, error)
	Update(ctx context.Context, order *model.TemplateM) error
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.TemplateM, error)
	Delete(ctx context.Context, id int64) error
}

// ConfigurationStore defines the interface for managing template configuration data storage.
type ConfigurationStore interface {
	Create(ctx context.Context, order *model.ConfigurationM) error
	CreateBatch(ctx context.Context, orders []*model.ConfigurationM) error
	Get(ctx context.Context, orderID string) (*model.ConfigurationM, error)
	Update(ctx context.Context, order *model.ConfigurationM) error
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.ConfigurationM, error)
	Delete(ctx context.Context, id int64) error
}

// HistoryStore defines the interface for managing message send history data storage.
type HistoryStore interface {
	Create(ctx context.Context, order *model.HistoryM) error
	Get(ctx context.Context, orderID string) (*model.HistoryM, error)
	Update(ctx context.Context, order *model.HistoryM) error
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.HistoryM, error)
	Delete(ctx context.Context, id int64) error
}

// InteractionStore defines the interface for managing uplink message data storage.
type InteractionStore interface {
	Create(ctx context.Context, order *model.InteractionM) error
	Get(ctx context.Context, orderID string) (*model.InteractionM, error)
	Update(ctx context.Context, order *model.InteractionM) error
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.InteractionM, error)
	Delete(ctx context.Context, id int64) error
}
