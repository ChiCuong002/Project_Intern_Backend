package helper

type CategoriesDropDown struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (CategoriesDropDown) TableName() string {
	return "categories"
}

type Categories struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	IsActive     bool   `json:"is_active"`
}

func (Categories) TableName() string {
	return "categories"
}
