package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginReset struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	Token     string
	Active    bool
	Notes     []*Note
	LoginID   uint
}

func (loginReset *LoginReset) Save(db *gorm.DB) error {
	return db.Save(loginReset).Error
}
