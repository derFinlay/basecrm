package models

import (
	"time"

	"gorm.io/gorm"
)

type BillingAddress struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
	City        string
	Street      string
	ZIP         string
	Number      string
	Notes       []*Note
	CustomerID  uint
	CreatedByID uint
	CreatedBy   *User
}

func (billingAddress *BillingAddress) Save(db *gorm.DB) error {
	return db.Save(billingAddress).Error
}
