package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (t *TemplateM) BeforeCreate(tx *gorm.DB) (err error) {
	if t.TemplateCode == "" {
		// Generate a new UUID for templateCode.
		t.TemplateCode = uuid.New().String()
	}

	return nil
}

// AfterCreate
//func (t *TemplateM) AfterCreate(tx *gorm.DB) (err error) {
//	return tx.Save(t).Error
//}
