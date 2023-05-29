package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	ProductID uint      `json:"product_id"`
	Product   *Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product,omitempty"`
	UserID    uint      `json:"user_id"`
	User      *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
