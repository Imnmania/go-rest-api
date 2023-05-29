package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	ProductRefer uint           `json:"product_id"`
	Product      *Product       `gorm:"foreignkey:ProductRefer"`
	UserRefer    uint           `json:"user_id"`
	User         *User          `gorm:"foreignkey:UserRefer"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
