package entities

import (
	"time"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Unit      string    `json:"unit"`
	Status    string    `json:"status"`
	ShopId    string    `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// change the table name to product since default one is products
func (Product) TableName() string {
	return "product"
}
