package mysql

import (
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
