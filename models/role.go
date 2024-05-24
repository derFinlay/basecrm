package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	Name      string
	Notes     []*Note
}

func (role *Role) Save(db *gorm.DB) error {
	return db.Save(role).Error
}
