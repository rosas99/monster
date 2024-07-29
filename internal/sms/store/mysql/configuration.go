package mysql

import (
	"context"
	"errors"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"gorm.io/gorm"
	"sort"
)

type configurationStore struct {
	db *gorm.DB
}

var _ store.ConfigurationStore = (*configurationStore)(nil)

// newConfigurations initializes a new configurationStore instance using the provided datastore.
func newConfigurations(db *gorm.DB) *configurationStore {
	return &configurationStore{db: db}
}

// Create adds a new configuration record in the datastore.
func (t *configurationStore) Create(ctx context.Context, template *model.ConfigurationM) error {
	return t.db.Create(&template).Error
}

// CreateBatch adds configuration records in the datastore.
func (t *configurationStore) CreateBatch(ctx context.Context, templates []*model.ConfigurationM) error {
	return t.db.Create(&templates).Error
}

// Delete removes a configuration record from the datastore.
func (t *configurationStore) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.TemplateM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// Get retrieves a configuration record from the datastore.
func (t *configurationStore) Get(ctx context.Context, id string) (*model.ConfigurationM, error) {
	var template model.ConfigurationM
	if err := t.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

// Update modifies an existing configuration record in the datastore.
func (t *configurationStore) Update(ctx context.Context, template *model.ConfigurationM) error {
	return t.db.Save(&template).Error
}

// List returns a list of configuration records that match the specified query conditions.
// It returns the total count of records and a slice of configuration records.
// The query dynamically applies filters, offset, limit, and order, based on provided list options.
func (t *configurationStore) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.ConfigurationM, err error) {
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

	// 升序排序
	sort.Sort(ByOrder(ret))
	//sort.SliceStable(cfgList, func(i, j int) bool {
	//	return cfgList[i].Order < cfgList[j].Order
	//})

	return count, ret, ans.Error
}
