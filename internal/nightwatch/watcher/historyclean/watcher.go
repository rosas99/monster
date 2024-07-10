// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

// Package secretsclean is a watcher implement used to delete expired keys from the database.
package historyclean

import (
	"context"
	"github.com/rosas99/monster/internal/nightwatch/watcher"
	"github.com/rosas99/monster/internal/pkg/client/store"
	"github.com/rosas99/monster/pkg/log"
	"sigs.k8s.io/cluster-api/util/secret"
	"time"
)

var _ watcher.Watcher = (*secretsCleanWatcher)(nil)

// watcher implement.
type secretsCleanWatcher struct {
	store store.Interface
}

// Run runs the watcher.
func (w *secretsCleanWatcher) Run() {
	_, histories, err := w.store.Sms().Histories().List(context.Background(), "")
	if err != nil {
		log.Errorw(err, "Failed to list secrets")
		return
	}

	for _, history := range histories {
		// 删除超过一年的历史记录
		if history.CreatedAt.Unix() < time.Now().AddDate(-1, 0, 0).Unix() {
			err := w.store.Sms().Histories().Delete(context.TODO(), history.ID)
			if err != nil {
				log.Warnw("Failed to delete secret from database", "userID", history.ID, "name", secret.Name)
				continue
			}
			log.Infow("Successfully deleted secret from database", "userID", history.ID, "name", secret.Name)
		}
	}
}

// Init initializes the watcher for later execution.
func (w *secretsCleanWatcher) Init(ctx context.Context, config *watcher.Config) error {
	w.store = config.Store
	return nil
}

func init() {
	watcher.Register("secretsclean", &secretsCleanWatcher{})
}
