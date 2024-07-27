package mysql

import (
	"sync"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/usercenter/store"
)

var (
	once sync.Once
	// Global variable that holds an already initialized *Datastore instance.
	S *Datastore
)

// Datastore is a concrete implementation of the IStore interface.
type Datastore struct {
	db *gorm.DB
}

// Ensures that Datastore implements the store.IStore interface.
var _ store.IStore = (*Datastore)(nil)

// NewStore creates an instance of type store.IStore.
func NewStore(db *gorm.DB) *Datastore {
	// Ensures that 'S' is only initialized once.
	once.Do(func() {
		S = &Datastore{db}
	})

	return S
}

// DB returns the *gorm.DB stored within the datastore.
func (ds *Datastore) DB() *gorm.DB {
	return ds.db
}

// Users returns a store.UserStore instance wrapping the database operations for users.
func (ds *Datastore) Users() store.UserStore {
	return newUsers(ds.db)
}
