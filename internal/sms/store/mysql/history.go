package mysql

import (
	"context"
	"errors"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"gorm.io/gorm"
)

// historyStore is an implementation of the HistoryStore interface
// that manages the interaction model in a datastore.
type historyStore struct {
	db *gorm.DB
}

var _ store.HistoryStore = (*historyStore)(nil)

// newHistories initializes a new historyStore instance using the provided datastore.
func newHistories(db *gorm.DB) *historyStore {
	return &historyStore{db: db}
}

// Create adds a new history record in the datastore.
func (t *historyStore) Create(ctx context.Context, template *model.HistoryM) error {
	return t.db.Create(&template).Error
}

// Delete removes a history record from the datastore.
func (t *historyStore) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.HistoryM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// Get retrieves a history record from the datastore.
func (t *historyStore) Get(ctx context.Context, templateCode string) (*model.HistoryM, error) {
	var template model.HistoryM
	if err := t.db.Where("template_code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// Update modifies an existing history record in the datastore.
func (t *historyStore) Update(ctx context.Context, template *model.HistoryM) error {
	return t.db.Save(&template).Error
}

// List returns a list of history records that match the specified query conditions.
// It returns the total count of records and a slice of history records.
// The query dynamically applies filters, offset, limit, and order, based on provided list options.
func (t *historyStore) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.HistoryM, err error) {

	options := meta.NewListOptions(opts...)
	if templateCode != "" {
		options.Filters["template_code"] = templateCode
	}
	// todo 对比 ucenter
	ans := t.db.
		Where(options.Filters).
		Offset(options.Offset).
		Limit(options.Limit).
		Order("id desc").
		Find(ret).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}
