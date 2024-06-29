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
	"gorm.io/gorm"
)

type hitories struct {
	db *gorm.DB
}

func newHistories(db *gorm.DB) *hitories {
	return &hitories{db: db}
}

func (t *hitories) Create(ctx context.Context, template *model.HistoryM) error {
	return t.db.Create(&template).Error
}

func (t *hitories) Get(ctx context.Context, templateCode string) (*model.HistoryM, error) {
	var template model.HistoryM
	if err := t.db.Where("template_code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (t *hitories) Update(ctx context.Context, template *model.HistoryM) error {
	return t.db.Save(&template).Error
}
func (t *hitories) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.HistoryM, err error) {

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
func (t *hitories) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.HistoryM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
