package user

import (
	"github.com/looplab/fsm"

	known "github.com/rosas99/monster/internal/pkg/known/usercenter"
)

func NewFSM(initial string, w *userWatcher) *fsm.FSM {
	return fsm.NewFSM(
		initial,
		fsm.Events{
			{Name: known.UserStatusRegistered, Src: []string{known.UserStatusRegistered}, Dst: known.UserStatusActive},
			{Name: known.UserStatusBlacklisted, Src: []string{known.UserStatusBlacklisted}, Dst: known.UserStatusDisabled},
			//{Name: known.UserStatusDisabled, Src: []string{known.UserStatusDisabled}, Dst: known.UserStatusDeleted},
			// todo 定义恢复用户激活状态
			{Name: known.UserStatusDisabled, Src: []string{known.UserStatusDisabled}, Dst: known.UserStatusActive},
		},
		fsm.Callbacks{
			known.UserStatusActive:   NewActiveUserCallback(w.store),
			known.UserStatusDisabled: NewDisableUserCallback(w.store),
			known.UserStatusDeleted:  NewDeleteUserCallback(w.store),
			// log, alert, save to store, etc for all events.
			// Alert the status of each step of the operation.
			UserEventAfterEvent: NewUserEventAfterEvent(w.store),
		},
	)
}
