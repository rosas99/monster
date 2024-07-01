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
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"gorm.io/gorm"
)

type configurations struct {
	db *gorm.DB
}

var _ store.ConfigurationStore = (*configurations)(nil)

func newConfigurations(db *gorm.DB) *configurations {
	return &configurations{db: db}
}

func (t *configurations) Create(ctx context.Context, template *model.ConfigurationM) error {
	return t.db.Create(&template).Error
}

func (t *configurations) CreateBatch(ctx context.Context, templates []*model.ConfigurationM) error {
	return t.db.Create(&templates).Error
}

func (t *configurations) Get(ctx context.Context, id string) (*model.ConfigurationM, error) {
	var template model.ConfigurationM
	if err := t.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (t *configurations) Update(ctx context.Context, template *model.ConfigurationM) error {
	return t.db.Save(&template).Error
}
func (t *configurations) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.ConfigurationM, err error) {
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
func (t *configurations) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.TemplateM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
