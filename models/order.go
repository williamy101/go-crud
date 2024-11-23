package models

import "time"

type Order struct {
	OrderID   int       `gorm:"primaryKey" json:"orderId"`
	ProductID int       `gorm:"index;not null" json:"productID"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	OrderData time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"orderDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ProductID" json:"product"`
}
