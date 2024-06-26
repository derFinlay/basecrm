package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID         uint        `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time   `gorm:"index" json:"createdAt"`
	UpdatedAt  time.Time   `gorm:"index" json:"updatedAt"`
	Positions  []*Position `json:"positions"`
	CustomerID uint        `json:"customerId"`
	Customer   *Customer   `json:"customer"`
	OrderID    uint        `json:"orderId"`
	Order      *Order      `json:"order"`
	Notes      []*Note     `json:"notes"`
}

func (invoice *Invoice) Save(db *gorm.DB) error {
	return db.Save(invoice).Error
}

func (invoice *Invoice) IsEmpty() bool {
	if invoice == nil {
		return true
	}
	return invoice.ID == 0
}
