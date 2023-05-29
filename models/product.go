package models

import (
	"time"
)

type Product struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Name         string    `gorm:"index" json:"name"`
	SerialNumber string    `gorm:"index" json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
