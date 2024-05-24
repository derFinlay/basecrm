package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time `gorm:"index"`
	Name        string
	Description string
	Price       float64
	Notes       []*Note
}

func (product *Product) Save(db *gorm.DB) error {
	return db.Save(product).Error
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
