// Package user is a watcher implement.
package user

import (
	"context"
	"github.com/rosas99/monster/internal/pkg/monsterx"

	"github.com/gammazero/workerpool"
	"github.com/looplab/fsm"

	"github.com/rosas99/monster/internal/nightwatch/watcher"
	"github.com/rosas99/monster/internal/pkg/client/store"
	known "github.com/rosas99/monster/internal/pkg/known/usercenter"
	"github.com/rosas99/monster/internal/usercenter/model"
	"github.com/rosas99/monster/pkg/log"
	stringsutil "github.com/rosas99/monster/pkg/util/strings"
)

var _ watcher.Watcher = (*userWatcher)(nil)

// watcher implement.
type userWatcher struct {
	store      store.Interface
	maxWorkers int64
}

type User struct {
	*model.UserM
	*fsm.FSM
}

// Run runs the watcher.
func (w *userWatcher) Run() {
	_, users, err := w.store.UserCenter().Users().List(context.Background())
	if err != nil {
		log.Errorw(err, "Failed to list users")
		return
	}

	allowOperations := []string{
		known.UserStatusRegistered,
		known.UserStatusBlacklisted,
		known.UserStatusDisabled,
	}

	wp := workerpool.New(int(w.maxWorkers))
	for _, user := range users {
		if !stringsutil.StringIn(user.Nickname, allowOperations) {
			continue
		}

		wp.Submit(func() {
			ctx := monsterx.NewUserM(context.Background(), user)
			u := &User{UserM: user, FSM: NewFSM(user.Nickname, w)}
			if err := u.Event(ctx, user.Status); err != nil {
				log.Errorw(err, "Failed to event user", "username", user.Username, "status", user.Nickname)
				return
			}

			return
		})
	}

	wp.StopWait()
}

// Init initializes the watcher for later execution.
func (w *userWatcher) Init(ctx context.Context, config *watcher.Config) error {
	w.store = config.Store
	w.maxWorkers = config.UserWatcherMaxWorkers
	return nil
}

func init() {
	watcher.Register("user", &userWatcher{})
}
