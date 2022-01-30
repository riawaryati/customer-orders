package models

type Item struct {
	// gorm.Model
	ItemID      int    `json:"lineItemId" gorm:"primaryKey;autoIncrement"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}
