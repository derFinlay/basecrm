package models

import (
	"time"

	"github.com/derfinlay/basecrm/service"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time `gorm:"index"`
	Name      string
	Username  string
	Password  string
	LastLogin time.Time
	Sessions  []*UserSession
	Notes     []*Note
}

func (user *User) Save(db *gorm.DB) error {
	return db.Save(user).Error
}

func (user *User) IsEmpty() bool {
	if user == nil {
		return true
	}
	return user.ID == 0
}

func (user *User) ResetPassword(newPassword string) error {
	hash, err := service.CreateHash(newPassword)
	if err != nil {
		return err
	}
	user.Password = hash
	return nil
}

func (user *User) CreateUserSession(db *gorm.DB) (*UserSession, error) {
	user.LastLogin = time.Now()
	user.Save(db)

	token := service.GenerateRandomString(32)

	userSession := &UserSession{
		Token:  token,
		Active: true,
		User:   user,
	}

	err := userSession.Save(db)

	return userSession, err
}
