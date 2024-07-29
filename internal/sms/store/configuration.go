// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package store

import (
	"context"
	"errors"
	"sort"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
)

// UserStore defines the interface for managing user data storage.
type ConfigurationStore interface {
	// Create adds a new user record to the database.
	Create(ctx context.Context, user *model.ConfigurationM) error
	// CreateBatch adds new users record to the database.
	CreateBatch(ctx context.Context, user []*model.ConfigurationM) error
	// List returns a slice of user records based on the specified query conditions.
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.ConfigurationM, error)
	// Get retrieves a user record by userID and username.
	Get(ctx context.Context, userID string, username string) (*model.ConfigurationM, error)
	// Update modifies an existing user record.
	Update(ctx context.Context, user *model.ConfigurationM) error
	// Delete removes a user record using the provided filters.
	Delete(ctx context.Context, filters map[string]any) error

	// Extensions
	// Fetch retrieves a user record using provided filters.
	Fetch(ctx context.Context, filters map[string]any) (*model.ConfigurationM, error)
	// GetByUsername retrieves a user record using username as the query condition.
	GetByUsername(ctx context.Context, username string) (*model.ConfigurationM, error)
}

// userStore is an implementation of the UserStore interface using a datastore.
type configurationStore struct {
	ds *datastore
}

// newUserStore returns a new instance of userStore with the provided datastore.
func newConfigurationStore(ds *datastore) *configurationStore {
	return &configurationStore{ds}
}

// db is an alias for d.ds.Core(ctx context.Context).
// It returns a pointer to a gorm.DB instance.
func (d *configurationStore) db(ctx context.Context) *gorm.DB {
	return d.ds.Core(ctx)
}

// Create adds a new user record to the database.
func (d *configurationStore) Create(ctx context.Context, user *model.ConfigurationM) error {
	return d.db(ctx).Create(&user).Error
}

// Create adds new users record to the database.
func (d *configurationStore) CreateBatch(ctx context.Context, user []*model.ConfigurationM) error {
	return d.db(ctx).Create(&user).Error
}

// List returns a slice of user records based on the specified query conditions
// along with the total number of records that match the given filters.
func (d *configurationStore) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.ConfigurationM, err error) {
	o := meta.NewListOptions(opts...)

	// List secrets for all users by default.
	if templateCode != "" {
		o.Filters["template_code"] = templateCode
	}

	ans := d.db(ctx).
		Where(o.Filters).
		Offset(o.Offset).
		Limit(o.Limit).
		Order("id desc").
		Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count)

	// 升序排序
	sort.Sort(ByOrder(ret))
	//sort.SliceStable(cfgList, func(i, j int) bool {
	//	return cfgList[i].Order < cfgList[j].Order
	//})

	return count, ret, ans.Error
}

// Fetch retrieves a user record from the database using the provided filters.
func (d *configurationStore) Fetch(ctx context.Context, filters map[string]any) (*model.ConfigurationM, error) {
	user := &model.ConfigurationM{}
	if err := d.db(ctx).Where(filters).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Get retrieves a user record by userID and username.
func (d *configurationStore) Get(ctx context.Context, userID string, username string) (*model.ConfigurationM, error) {
	return d.Fetch(ctx, map[string]any{"user_id": userID, "username": username})
}

// GetByUsername retrieves a user record using the provided username.
func (d *configurationStore) GetByUsername(ctx context.Context, username string) (*model.ConfigurationM, error) {
	return d.Fetch(ctx, map[string]any{"username": username})
}

// Update modifies an existing user record in the database.
func (d *configurationStore) Update(ctx context.Context, user *model.ConfigurationM) error {
	return d.db(ctx).Save(user).Error
}

// Delete removes a user record from the database using the provided filters.
// It returns an error if the deletion process encounters an issue other than a missing record.
func (d *configurationStore) Delete(ctx context.Context, filters map[string]any) error {
	err := d.db(ctx).Where(filters).Delete(&model.ConfigurationM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
