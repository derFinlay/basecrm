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
	UserID            *uint
	CustomerID        *uint
	BillingAddressID  *uint
	DeliveryAddressID *uint
	InvoiceID         *uint
	LoginID           *uint
	LoginResetID      *uint
	CreatedByID       *uint
	OrderID           *uint
	PositionID        *uint
	ProductID         *uint
	RoleID            *uint
	TelID             *uint
	CreatedBy         *User
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
