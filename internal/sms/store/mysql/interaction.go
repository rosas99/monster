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

type interactions struct {
	db *gorm.DB
}

var _ store.InteractionStore = (*interactions)(nil)

func newInteractions(db *gorm.DB) *interactions {
	return &interactions{db: db}
}

func (t *interactions) Create(ctx context.Context, template *model.InteractionM) error {
	return t.db.Create(&template).Error
}

func (t *interactions) CreateBatch(ctx context.Context, templates []*model.InteractionM) error {
	return t.db.Create(&templates).Error
}

func (t *interactions) Get(ctx context.Context, id string) (*model.InteractionM, error) {
	var template model.InteractionM
	if err := t.db.Where("id = ?", id).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (t *interactions) Update(ctx context.Context, template *model.InteractionM) error {
	return t.db.Save(&template).Error
}
func (t *interactions) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.InteractionM, err error) {
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
func (t *interactions) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.InteractionM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
