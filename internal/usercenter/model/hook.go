package model

import (
	"github.com/rosas99/monster/pkg/auth"
	"gorm.io/gorm"
)

func (t *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	//if t.Username == "" {
	//	// Generate a new UUID for SecretKey.
	//	t.Username = uuid.New().String()
	//}

	// Encrypt the user password.
	t.Password, err = auth.Encrypt(t.Password)
	if err != nil {
		return err
	}
	// 可以考虑zid的生成方式
	return nil
}

// AfterCreate todo 待确认 生成带盐值的随机id 非主键
func (t *UserM) AfterCreate(tx *gorm.DB) (err error) {
	//t.Username = zid.Template.New(uint64(t.ID)) // Generate and set a new order ID.

	return tx.Save(t).Error // Save the updated order record to the database.
}
