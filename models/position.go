package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
	Name        string
	Description string
	UnitPrice   float64
	Amount      int
	InvoiceID   uint
	Notes       []*Note
}

func (position *Position) Save(db *gorm.DB) error {
	return db.Save(position).Error
}
