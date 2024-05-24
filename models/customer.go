package models

import (
	"time"

	"github.com/derfinlay/basecrm/service"
	"gorm.io/gorm"
)

type Customer struct {
	ID                uint      `gorm:"primarykey"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time `gorm:"index"`
	Name              string
	Gender            string
	Orders            []*Order
	BillingAddresses  []*BillingAddress
	DeliveryAddresses []*DeliveryAddress
	Tels              []*Tel
	CreatedByID       uint
	CreatedBy         *User
	Notes             []*Note
	LoginId           uint
	Login             *Login
}

func (customer *Customer) Save(db *gorm.DB) error {
	return db.Save(customer).Error
}

func (customer *Customer) CreateLogin(email string, emailClient *service.EmailClient, db *gorm.DB) error {
	err := service.ValidateEmail(email)
	if err != nil {
		return err
	}
	defaultPassword := service.GenerateRandomString(12)
	hash, err := service.CreateHash(defaultPassword)
	if err != nil {
		return err
	}
	login := &Login{
		Email:    email,
		Password: hash,
	}

	customer.Login = login

	login.SendDefaultPasswordEmail(defaultPassword, emailClient)

	return customer.Save(db)
}
