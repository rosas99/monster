package mysql

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/sms/store"
)

// Singleton instance variables.
var (
	once sync.Once
	s    *Datastore
)

// Datastore is an implementation of IStore that provides methods
// to perform operations on a database using gorm library.
type Datastore struct {
	db *gorm.DB
}

// Ensure datastore implements IStore.
var _ store.IStore = (*Datastore)(nil)

// NewStore initializes a new datastore instance using the provided DB gorm instance.
// It also creates a singleton instance for the datastore.
func NewStore(db *gorm.DB) *Datastore {
	once.Do(func() {
		s = &Datastore{db}
	})
	return s
}

// Templates returns an initialized instance of TemplateStore.
func (ds *Datastore) Templates() store.TemplateStore {
	return newTemplates(ds.db)
}

// Configurations returns an initialized instance of ConfigurationStore.
func (ds *Datastore) Configurations() store.ConfigurationStore {
	return newConfigurations(ds.db)
}

// Histories returns an initialized instance of HistoryStore.
func (ds *Datastore) Histories() store.HistoryStore {
	return newHistories(ds.db)
}

// Interactions returns an initialized instance of InteractionStore.
func (ds *Datastore) Interactions() store.InteractionStore {
	return newInteractions(ds.db)
}

// transactionKey is an unique key used in context to store
// transaction instances to be shared between multiple operations.
type transactionKey struct{}

// Core retrieves the current transactional DB instance if it exists
// in context or falls back to the main database.
func (ds *Datastore) Core(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(transactionKey{}).(*gorm.DB)
	if ok {
		return tx
	}

	return ds.db
}

// TX starts a transaction using the main DB context
// and passes the transactional context to the provided function.
func (ds *Datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}
