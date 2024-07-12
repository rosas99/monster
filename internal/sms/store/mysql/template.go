package mysql

import (
	"context"
	"errors"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"gorm.io/gorm"
)

// templateStore is an implementation of the TemplateStore interface
// that manages the template model in a datastore.
type templateStore struct {
	db *gorm.DB
}

var _ store.TemplateStore = (*templateStore)(nil)

// newTemplates initializes a new templateStore instance using the provided datastore.
func newTemplates(db *gorm.DB) *templateStore {
	return &templateStore{db: db}
}

// Create adds a new template record in the datastore.
func (t *templateStore) Create(ctx context.Context, template *model.TemplateM) error {
	return t.db.Create(&template).Error
}

// Delete removes a template record from the datastore.
func (t *templateStore) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.TemplateM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// Get retrieves a template record from the datastore.
func (t *templateStore) Get(ctx context.Context, templateCode string) (*model.TemplateM, error) {
	var template model.TemplateM
	if err := t.db.Where("template_code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// Update modifies an existing template record in the datastore.
func (t *templateStore) Update(ctx context.Context, template *model.TemplateM) error {
	return t.db.Save(&template).Error
}

// List returns a list of template records that match the specified query conditions.
// It returns the total count of records and a slice of template records.
// The query dynamically applies filters, offset, limit, and order, based on provided list options.
func (t *templateStore) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.TemplateM, err error) {

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
