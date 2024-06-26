package models

import (
	"time"

	"github.com/derfinlay/basecrm/service"
	"gorm.io/gorm"
)

type Customer struct {
	ID                uint               `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time          `gorm:"index" json:"createdAt"`
	UpdatedAt         time.Time          `gorm:"index" json:"updatedAt"`
	Name              string             `json:"name"`
	Gender            string             `json:"gender"`
	Orders            []*Order           `json:"orders"`
	BillingAddress    *BillingAddress    `json:"billingAddress"`
	DeliveryAddresses []*DeliveryAddress `json:"deliveryAddresses"`
	Tels              []*Tel             `json:"tels"`
	CreatedByID       uint               `json:"createdById"`
	CreatedBy         *User              `json:"createdBy"`
	Notes             []*Note            `json:"notes"`
	LoginId           *uint              `json:"loginId"`
	Login             *Login             `json:"login"`
}

func (customer *Customer) Save(db *gorm.DB) error {
	return db.Save(customer).Error
}

func (customer *Customer) IsEmpty() bool {
	if customer == nil {
		return true
	}
	return customer.ID == 0
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
