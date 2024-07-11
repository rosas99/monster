package mysql

import (
	"context"
	"errors"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"gorm.io/gorm"
)

// interactionStore is an implementation of the InteractionStore interface
// that manages the interaction model in a datastore.
type interactionStore struct {
	db *gorm.DB
}

var _ store.InteractionStore = (*interactionStore)(nil)

// newInteractions initializes a new interactionStore instance using the provided datastore.
func newInteractions(db *gorm.DB) *interactionStore {
	return &interactionStore{db: db}
}

// Create adds a new interaction record in the datastore.
func (t *interactionStore) Create(ctx context.Context, template *model.InteractionM) error {
	return t.db.Create(&template).Error
}

// CreateBatch adds interaction records in the datastore.
func (t *interactionStore) CreateBatch(ctx context.Context, templates []*model.InteractionM) error {
	return t.db.Create(&templates).Error
}

// Delete removes an interaction record from the datastore.
func (t *interactionStore) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.InteractionM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// Get retrieves an interaction record from the datastore.
func (t *interactionStore) Get(ctx context.Context, id string) (*model.InteractionM, error) {
	var template model.InteractionM
	if err := t.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// Update modifies an existing interaction record in the datastore.
func (t *interactionStore) Update(ctx context.Context, template *model.InteractionM) error {
	return t.db.Save(&template).Error
}

// List returns a list of interaction records that match the specified query conditions.
// It returns the total count of records and a slice of interaction records.
// The query dynamically applies filters, offset, limit, and order, based on provided list options.
func (t *interactionStore) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.InteractionM, err error) {
	options := meta.NewListOptions(opts...)
	if templateCode != "" {
		options.Filters["template_code"] = templateCode
	}
	ans := t.db.
		Where(options.Filters).
		Offset(options.Offset).
		Limit(options.Limit).
		Order("id desc").
		Find(&ret).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}
