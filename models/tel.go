package models

import (
	"time"

	"gorm.io/gorm"
)

type Tel struct {
	ID                uint      `gorm:"primarykey"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time `gorm:"index"`
	Tel               string
	CustomerID        uint
	DeliveryAddressID uint

	Notes []*Note
}

func (tel *Tel) Save(db *gorm.DB) error {
	return db.Save(tel).Error
}
