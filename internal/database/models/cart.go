package models

type CartItem struct {
	Default
	SessionId string
	ProductId uint64 `json:"product_id" gorm:"not null"`
	Quantity  uint   `json:"quantity" gorm:"not null"`
}
