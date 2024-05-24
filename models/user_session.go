package models

import (
	"time"

	"gorm.io/gorm"
)

type UserSession struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	Token     string
	Active    bool
	UserId    uint
	User      *User
	UserID    uint
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
