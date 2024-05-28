package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"index" json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Notes       []*Note   `json:"notes"`
}

func (product *Product) Save(db *gorm.DB) error {
	return db.Save(product).Error
}

func (product *Product) IsEmpty() bool {
	if product == nil {
		return true
	}
	return product.ID == 0
}

func RemoveDuplicates[T string | int | *Product](input []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range input {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
