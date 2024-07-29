// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

package store

import "github.com/rosas99/monster/internal/sms/model"

const defaultLimitValue = 20

// defaultLimit 设置默认查询记录数.
func defaultLimit(limit int) int {
	if limit == 0 {
		limit = defaultLimitValue
	}

	return limit
}

type ByOrder []*model.ConfigurationM

func (o ByOrder) Len() int           { return len(o) }
func (o ByOrder) Less(i, j int) bool { return o[i].Order < o[j].Order }
func (o ByOrder) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
