package entities

type Category struct {
	CatId        int16      `json:"cat_id"`
	Name         string     `json:"name"`
	ParentCid    int16      `json:"parent_id"`
	CatLevel     int8       `json:"cat_level"`
	ShowStatus   int8       `json:"show_status"`
	Sort         int16      `json:"sort"`
	ProductCount float32    `json:"product_count"`
	Children     []Category `gorm:"-" json:"children"`
}

// change the table name to category since default one is categories
func (Category) TableName() string {
	return "category"
}
