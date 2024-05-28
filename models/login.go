package models

import (
	"time"

	"github.com/derfinlay/basecrm/service"
	"gorm.io/gorm"
)

type Login struct {
	ID          uint          `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time     `gorm:"index" json:"createdAt"`
	UpdatedAt   time.Time     `gorm:"index" json:"updatedAt"`
	Password    string        `json:"-"`
	Email       string        `json:"email"`
	LastLogin   time.Time     `json:"lastLogin"`
	LoginResets []*LoginReset `json:"resets"`
	Notes       []*Note       `json:"notes"`
}

func (login *Login) Save(db *gorm.DB) error {
	return db.Save(login).Error
}

func (login *Login) IsEmpty() bool {
	if login == nil {
		return true
	}
	return login.ID == 0
}

func (login *Login) SendDefaultPasswordEmail(password string, emailClient *service.EmailClient) error {
	return emailClient.SendEmail([]string{login.Email}, "Ihr Standartpasswort lautet: "+password)
}

func (login *Login) SendLoginResetEmail(reset *LoginReset, emailClient *service.EmailClient) error {
	return emailClient.SendEmail([]string{login.Email}, "Ihr Passwort Reset Link: "+reset.Token)
}

func (login *Login) ResetPassword(newPassword string, db *gorm.DB) error {
	hash, err := service.CreateHash(newPassword)
	if err != nil {
		return err
	}
	login.Password = hash

	return login.Save(db)
}
