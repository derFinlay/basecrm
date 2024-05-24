package models

import (
	"time"

	"github.com/derfinlay/basecrm/service"
	"gorm.io/gorm"
)

type Login struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
	Password    string
	Email       string
	LastLogin   time.Time
	LoginResets []*LoginReset
	Notes       []*Note
}

func (login *Login) Save(db *gorm.DB) error {
	return db.Save(login).Error
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
