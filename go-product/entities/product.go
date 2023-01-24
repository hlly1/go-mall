package entities

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
}

// change the table name to stock since default one is stocks
func (Product) TableName() string {
	return "product"
}
