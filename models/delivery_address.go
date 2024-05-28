package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryAddress struct {
	ID          uint      `gorm:"primarykey" json:"id`
	CreatedAt   time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index" json:"updatedAt"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	ZIP         string    `json:"zip"`
	Number      string    `json:"number"`
	CustomerID  uint      `json:"customerId"`
	Tels        []*Tel    `json:"tels"`
	Notes       []*Note   `json:"notes"`
	CreatedByID uint      `json:"createdById"`
	CreatedBy   *User     `json:"createdBy"`
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
