package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ProductID   int             `gorm:"primaryKey" json:"productId"`
	Name        string          `gorm:"not null" json:"name"`
	Description *string         `json:"description"`
	Price       decimal.Decimal `gorm:"not null" json:"price"`
	Category    *string         `json:"category"`
	ImagePath   *string         `json:"imagePath"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}
