package models

import (
	"time"

	"gorm.io/gorm"
)

type BillingAddress struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index" json:"updatedAt"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	ZIP         string    `json:"zip"`
	Number      string    `json:"number"`
	CustomerID  uint      `json:"customerId"`
	Customer    *Customer `json:"customer"`
	CreatedByID uint      `json:"createdById"`
	CreatedBy   *User     `json:"createdBy"`
}

func (billingAddress *BillingAddress) Save(db *gorm.DB) error {
	return db.Save(billingAddress).Error
}

func (billingAddress *BillingAddress) IsEmpty() bool {
	if billingAddress == nil {
		return true
	}
	return billingAddress.ID == 0
}
