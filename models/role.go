package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"index" json:"updatedAt"`
	Name      string    `json:"name"`
	Notes     []*Note   `json:"notes"`
}

func (role *Role) Save(db *gorm.DB) error {
	return db.Save(role).Error
}

func (role *Role) IsEmpty() bool {
	if role == nil {
		return true
	}
	return role.ID == 0
}
