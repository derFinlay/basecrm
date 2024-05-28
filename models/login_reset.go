package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginReset struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"index" json:"updatedAt"`
	Token     string    `json:"token"`
	Active    bool      `json:"active"`
	Notes     []*Note   `json:"notes"`
	LoginID   uint      `json:"loginId"`
	Login     *Login    `json:"login"`
}

func (loginReset *LoginReset) Save(db *gorm.DB) error {
	return db.Save(loginReset).Error
}

func (loginReset *LoginReset) IsEmpty() bool {
	if loginReset == nil {
		return true
	}
	return loginReset.ID == 0
}
