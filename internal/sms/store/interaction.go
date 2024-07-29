// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package store

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/sms/model"
)

// UserStore defines the interface for managing user data storage.
type InteractionStore interface {
	// Create adds a new user record to the database.
	Create(ctx context.Context, user *model.InteractionM) error
	// List returns a slice of user records based on the specified query conditions.
	List(ctx context.Context, opts ...meta.ListOption) (int64, []*model.InteractionM, error)
	// Get retrieves a user record by userID and username.
	Get(ctx context.Context, userID string, username string) (*model.InteractionM, error)
	// Update modifies an existing user record.
	Update(ctx context.Context, user *model.InteractionM) error
	// Delete removes a user record using the provided filters.
	Delete(ctx context.Context, filters map[string]any) error

	// Extensions
	// Fetch retrieves a user record using provided filters.
	Fetch(ctx context.Context, filters map[string]any) (*model.InteractionM, error)
	// GetByUsername retrieves a user record using username as the query condition.
	GetByUsername(ctx context.Context, username string) (*model.InteractionM, error)
}

// userStore is an implementation of the UserStore interface using a datastore.
type interactionStore struct {
	ds *datastore
}

// newUserStore returns a new instance of userStore with the provided datastore.
func newInteractionStore(ds *datastore) *interactionStore {
	return &interactionStore{ds}
}

// db is an alias for d.ds.Core(ctx context.Context).
// It returns a pointer to a gorm.DB instance.
func (d *interactionStore) db(ctx context.Context) *gorm.DB {
	return d.ds.Core(ctx)
}

// Create adds a new user record to the database.
func (d *interactionStore) Create(ctx context.Context, user *model.InteractionM) error {
	return d.db(ctx).Create(&user).Error
}

// List returns a slice of user records based on the specified query conditions
// along with the total number of records that match the given filters.
func (d *interactionStore) List(ctx context.Context, opts ...meta.ListOption) (count int64, ret []*model.InteractionM, err error) {
	o := meta.NewListOptions(opts...)

	ans := d.db(ctx).
		Where(o.Filters).
		Offset(o.Offset).
		Limit(o.Limit).
		Order("id desc").
		Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}

// Fetch retrieves a user record from the database using the provided filters.
func (d *interactionStore) Fetch(ctx context.Context, filters map[string]any) (*model.InteractionM, error) {
	user := &model.InteractionM{}
	if err := d.db(ctx).Where(filters).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Get retrieves a user record by userID and username.
func (d *interactionStore) Get(ctx context.Context, userID string, username string) (*model.InteractionM, error) {
	return d.Fetch(ctx, map[string]any{"user_id": userID, "username": username})
}

// GetByUsername retrieves a user record using the provided username.
func (d *interactionStore) GetByUsername(ctx context.Context, username string) (*model.InteractionM, error) {
	return d.Fetch(ctx, map[string]any{"username": username})
}

// Update modifies an existing user record in the database.
func (d *interactionStore) Update(ctx context.Context, user *model.InteractionM) error {
	return d.db(ctx).Save(user).Error
}

// Delete removes a user record from the database using the provided filters.
// It returns an error if the deletion process encounters an issue other than a missing record.
func (d *interactionStore) Delete(ctx context.Context, filters map[string]any) error {
	err := d.db(ctx).Where(filters).Delete(&model.InteractionM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
