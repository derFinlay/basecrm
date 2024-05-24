package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID                uint      `gorm:"primarykey"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time `gorm:"index"`
	Title             string
	Content           string
	UserID            uint
	CustomerID        uint
	BillingAddressID  uint
	DeliveryAddressID uint
	InvoiceID         uint
	LoginID           uint
	LoginResetID      uint
	CreatedByID       uint
	OrderID           uint
	PositionID        uint
	ProductID         uint
	RoleID            uint
	TelID             uint
	CreatedBy         *User
}

func (note *Note) Save(db *gorm.DB) error {
	return db.Save(note).Error
}
