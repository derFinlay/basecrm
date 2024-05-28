package models

import (
	"time"

	"gorm.io/gorm"
)

type Tel struct {
	ID                uint      `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"index" json:"updatedAt"`
	Tel               string    `json:"tel"`
	CustomerID        *uint     `json:"-"`
	DeliveryAddressID *uint     `json:"-"`
	Notes             []*Note   `json:"notes"`
}

func (tel *Tel) Save(db *gorm.DB) error {
	return db.Save(tel).Error
}

func (tel *Tel) IsEmpty() bool {
	if tel == nil {
		return true
	}
	return tel.ID == 0
}
