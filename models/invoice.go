package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time `gorm:"index"`
	Positions  []*Position
	CustomerID uint
	Customer   *Customer
	OrderID    uint
	Order      *Order
	Notes      []*Note
}

func (invoice *Invoice) Save(db *gorm.DB) error {
	return db.Save(invoice).Error
}
