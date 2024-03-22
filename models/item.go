package models

type Item struct {
	ItemID      int    `json:"item_id" gorm:"primary_key"`
	ItemCode    string `json:"item_code" gorm:"unique;not null; type:varchar(100)"`
	Description string `json:"description" gorm:"text"`
	Quantity    int    `json:"quantity" gorm:"int"`
	OrderID     int    `json:"order_id"`
}