package store

//go:generate mockgen -destination mock_store.go -package store github.com/rosas99/monster/internal/fakeserver/store IStore,OrderStore

import (
	"context"
	"sync"

	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/usercenter/model"
)

// IStore defines the methods that the Store layer needs to implement.
type IStore interface {
	Users() UserStore
}

// UserStore defines the methods implemented by the store layer for the user module.
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
	Update(ctx context.Context, user *model.UserM) error
	List(ctx context.Context, opts ...meta.ListOption) (count int64, ret []*model.UserM, err error)
	Delete(ctx context.Context, id int64) error
}

var (
	once sync.Once
	// S Global variable that allows other packages to directly call the already initialized S instance.
	S IStore
)

// SetStore set the onex-fakeserver store instance in a global variable `S`.
// Direct use the global `S` is not recommended as this may make dependencies and calls unclear.
func SetStore(store IStore) {
	once.Do(func() {
		S = store
	})
}
