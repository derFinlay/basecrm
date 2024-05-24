package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryAddress struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
	City        string
	Street      string
	ZIP         string
	Number      string
	CustomerID  uint
	TelID       uint
	Tel         []*Tel
	Notes       []*Note
	CreatedByID uint
	CreatedBy   *User
}

func (deliveryAddress *DeliveryAddress) Save(db *gorm.DB) error {
	return db.Save(deliveryAddress).Error
}

func (deliveryAddress *DeliveryAddress) IsEmpty() bool {
	if deliveryAddress == nil {
		return true
	}
	return deliveryAddress.ID == 0
}
