package user

import (
	"context"
	known "github.com/rosas99/monster/internal/pkg/known/usercenter"
	"github.com/rosas99/monster/internal/pkg/monsterx"
	"time"

	"github.com/looplab/fsm"

	"github.com/rosas99/monster/internal/pkg/client/store"
	"github.com/rosas99/monster/pkg/log"
)

const (
	UserEventAfterEvent = "after_event" // fsm default event.
)

func NewActiveUserCallback(store store.Interface) fsm.Callback {
	return func(ctx context.Context, event *fsm.Event) {
		userM := monsterx.FromUserM(ctx)
		log.Infow("Now active user", "event", event.Event, "username", userM.Username)
		// Fake active user operations.
		time.Sleep(5 * time.Second)
		log.Infow("Success to active user", "event", event.Event, "username", userM.Username)
	}
}

func NewDisableUserCallback(store store.Interface) fsm.Callback {
	return func(ctx context.Context, event *fsm.Event) {
		userM := monsterx.FromUserM(ctx)
		log.Infow("Now disable user", "event", event.Event, "username", userM.Username)
		// Fake disable user operations.
		time.Sleep(5 * time.Second)
		log.Infow("Success to disable user", "event", event.Event, "username", userM.Username)
	}
}

func NewDeleteUserCallback(store store.Interface) fsm.Callback {
	return func(ctx context.Context, event *fsm.Event) {
		userM := monsterx.FromUserM(ctx)
		log.Infow("Now delete user", "event", event.Event, "username", userM.Username)
		// Fake delete user operations.
		time.Sleep(5 * time.Second)
		log.Infow("Success to delete user", "event", event.Event, "username", userM.Username)
	}
}

func NewUserEventAfterEvent(store store.Interface) fsm.Callback {
	return func(ctx context.Context, event *fsm.Event) {
		alarmStatus := "success"
		userM := monsterx.FromUserM(ctx)

		defer func() {
			log.Infow("This is a fake alarm message", "status", alarmStatus, "username", userM.Username)
		}()

		if event.Err != nil {
			alarmStatus = "failed"
			log.Errorw(event.Err, "Failed to handle event", "event", event.Event)
			// We can add some alerts here in the future.
			return
		}

		user := monsterx.FromUserM(ctx)
		user.Username = event.FSM.Current()
		if err := store.UserCenter().Users().Update(ctx, user); err != nil {
			log.Errorw(err, "Failed to update status into database", "event", event.Event)
		}

		if user.Username == known.UserStatusDeleted {
			// We can add some lark card here in the future.
			log.Infow("Finish to handle user", "event", event.Event, "username", user.Username)
		}
	}
}
