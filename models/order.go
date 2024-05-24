package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                uint      `gorm:"primarykey"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time `gorm:"index"`
	tax               float64
	Due               time.Time
	PrintedAt         time.Time
	PayedAt           time.Time
	DoneAt            time.Time
	CustomerID        uint
	Customer          *Customer
	BillingAddressID  uint
	BillingAddress    *BillingAddress
	DeliveryAddressID uint
	DeliveryAddress   *DeliveryAddress
	Notes             []*Note
	Invoices          []*Invoice
	CreatedByID       uint
	CreatedBy         *User
}

func (order *Order) Save(db *gorm.DB) error {
	return db.Save(order).Error
}
