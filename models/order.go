package models

import (
	"time"
)

type Order struct {
	// gorm.Model
	OrderID      int       `json:"orderId" gorm:"primaryKey;autoIncrement"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;references:OrderID;"`
}

type OrderRequest struct {
	OrderID      int    `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item `json:"items"`
}
