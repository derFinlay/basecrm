package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID                uint      `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time `gorm:"index" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"index" json:"updatedAt"`
	Title             string    `json:"title"`
	Content           string    `json:"content"`
	UserID            *uint     `json:"-"`
	CustomerID        *uint     `json:"-"`
	DeliveryAddressID *uint     `json:"-"`
	InvoiceID         *uint     `json:"-"`
	LoginID           *uint     `json:"-"`
	LoginResetID      *uint     `json:"-"`
	CreatedByID       *uint     `json:"-"`
	OrderID           *uint     `json:"-"`
	PositionID        *uint     `json:"-"`
	ProductID         *uint     `json:"-"`
	RoleID            *uint     `json:"-"`
	TelID             *uint     `json:"-"`
	CreatedBy         *User     `json:"user"`
}

func (note *Note) Save(db *gorm.DB) error {
	return db.Save(note).Error
}

func (note *Note) IsEmpty() bool {
	if note == nil || (note.Content == "" && note.Title == "") {
		return true
	}
	return note.ID == 0
}
