package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index" json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UnitPrice   float32   `json:"unitPrice"`
	Amount      int       `json:"amount"`
	InvoiceID   uint      `json:"invoiceID"`
	Invoice     *Invoice  `json:"invoice"`
	Notes       []*Note   `json:"notes"`
}

func (position *Position) Save(db *gorm.DB) error {
	return db.Save(position).Error
}

func (position *Position) IsEmpty() bool {
	if position == nil {
		return true
	}
	return position.ID == 0
}
