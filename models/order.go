package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                uint             `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time        `gorm:"index" json:"createdAt"`
	UpdatedAt         time.Time        `gorm:"index" json:"updatedAt"`
	Tax               float32          `json:"tax"`
	Due               time.Time        `json:"due"`
	PrintedAt         time.Time        `json:"printedAt"`
	PayedAt           time.Time        `json:"paidAt"`
	DoneAt            time.Time        `json:"doneAt"`
	CustomerID        uint             `json:"customerId"`
	Customer          *Customer        `json:"customer"`
	BillingAddressID  uint             `json:"billingAddressId"`
	BillingAddress    *BillingAddress  `json:"billingAddress"`
	DeliveryAddressID uint             `json:"deliveryAddressId"`
	DeliveryAddress   *DeliveryAddress `json:"deliveryAddress"`
	Notes             []*Note          `json:"notes"`
	Invoices          []*Invoice       `json:"invoices"`
	CreatedByID       uint             `json:"createdById"`
	CreatedBy         *User            `json:"createdBy"`
}

func (order *Order) Save(db *gorm.DB) error {
	return db.Save(order).Error
}

func (order *Order) IsEmpty() bool {
	if order == nil {
		return true
	}
	return order.ID == 0
}
