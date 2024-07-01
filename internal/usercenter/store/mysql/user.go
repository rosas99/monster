// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package mysql

import (
	"context"
	"errors"
	"github.com/rosas99/monster/internal/pkg/meta"
	"github.com/rosas99/monster/internal/usercenter/model"
	"github.com/rosas99/monster/internal/usercenter/store"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

var _ store.UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

func (t *users) Create(ctx context.Context, template *model.UserM) error {
	return t.db.Create(&template).Error
}

func (t *users) Get(ctx context.Context, templateCode string) (*model.UserM, error) {
	var template model.UserM
	if err := t.db.Where("template_code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (t *users) Update(ctx context.Context, template *model.UserM) error {
	return t.db.Save(&template).Error
}
func (t *users) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.UserM, err error) {

	options := meta.NewListOptions(opts...)
	if templateCode != "" {
		options.Filters["template_code"] = templateCode
	}
	// todo 对比 user center
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
func (t *users) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.UserM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
