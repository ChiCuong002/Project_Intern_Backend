package helper

type ProductInsert struct {
	ProductID   uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	CategoryID  uint
	StatusID    uint
	ProductName string
	Description string
	Price       float64
	Quantity    uint
}

type ProductImageInsert struct {
	ProductImageID uint `gorm:"primaryKey;autoIncrement"`
	ProductID      uint
	ImageID        uint
}

type ImageInsert struct {
	ImageID   uint `gorm:"primaryKey;autoIncrement"`
	BucketKey string
	Path      string
}

func (ProductInsert) TableName() string {
	return "products"
}
func (ProductImageInsert) TableName() string {
	return "product_images"
}
func (ImageInsert) TableName() string {
	return "images"
}

type UpdateProduct struct {
	ProductID   uint    `json:"-"`
	CategoryID  uint    `json:"category_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
}

func (UpdateProduct) TableName() string {
	return "products"
}
