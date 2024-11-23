package models

import "time"

type Inventory struct {
	InventoryID int       `gorm:"primaryKey" json:"inventoryId"`
	ProductID   int       `gorm:"index;not null" json:"productId"`
	Stock       int       `gorm:"default:0;not null" json:"stock"`
	Location    string    `gorm:"size:100;not null" json:"location"`
	Product     Product   `gorm:"foreignKey:ProductID;references:ProductID" json:"product"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
