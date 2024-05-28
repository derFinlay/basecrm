package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt time.Time `gorm:"index" json:"updatedAt"`
	Token     string    `json:"-"`
	Active    bool      `json:"active"`
	User      *User     `json:"user,omitempty"`
	UserID    uint      `json:"userId,omitempty"`
}

func (session *UserSession) Save(db *gorm.DB) error {
	return db.Save(session).Error
}

func (session *UserSession) IsEmpty() bool {
	if session == nil {
		return true
	}
	return session.ID == 0
}

func (session *UserSession) Logout(db *gorm.DB) error {
	session.Active = false
	return session.Save(db)
}
