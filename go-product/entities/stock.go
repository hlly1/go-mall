package entities

// default table name is stocks
type Stock struct {
	ProductId string  `json:"product_id"`
	Stock     float64 `json:"stock"`
	StockUnit string  `json:"stock_unit"`
}

// change the table name to stock since default one is stocks
func (Stock) TableName() string {
	return "stock"
}
