package entities

// default table name is stocks
type Stock struct {
	ProductId		string
	Stock			float64
	Stock_unit		string
}

// change the table name to stock since default one is stocks
func (Stock) TableName() string{
	return "stock"
}