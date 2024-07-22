package model

import (
	"github.com/rosas99/monster/pkg/auth"
	"gorm.io/gorm"
)

func (t *UserM) BeforeCreate(tx *gorm.DB) (err error) {

	// Encrypt the user password.
	t.Password, err = auth.Encrypt(t.Password)
	if err != nil {
		return err
	}
	return nil
}
