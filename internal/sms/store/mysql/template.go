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

type TemplateStore interface {
	Create(ctx context.Context, order *model.TemplateM) error
	Get(ctx context.Context, templateCode string) (*model.TemplateM, error)
	Update(ctx context.Context, order *model.TemplateM) error
	List(ctx context.Context, templateCode string, opts ...meta.ListOption) (int64, []*model.TemplateM, error)
	Delete(ctx context.Context, id int64) error
}

type templates struct {
	db *gorm.DB
}

var _ store.TemplateStore = (*templates)(nil)

func newTemplates(db *gorm.DB) *templates {
	return &templates{db: db}
}

func (t *templates) Create(ctx context.Context, template *model.TemplateM) error {
	return t.db.Create(&template).Error
}

func (t *templates) Get(ctx context.Context, templateCode string) (*model.TemplateM, error) {
	var template model.TemplateM
	if err := t.db.Where("template_code = ?", templateCode).First(&template).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (t *templates) Update(ctx context.Context, template *model.TemplateM) error {
	return t.db.Save(&template).Error
}
func (t *templates) List(ctx context.Context, templateCode string, opts ...meta.ListOption) (count int64, ret []*model.TemplateM, err error) {

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
		Find(&ret).
		Limit(-1).
		Count(&count)

	return count, ret, ans.Error
}
func (t *templates) Delete(ctx context.Context, id int64) error {
	err := t.db.Where("id = ?", id).Delete(&model.TemplateM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
