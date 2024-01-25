package helper

type DetailProductRes struct {
	ProductID     uint           `json:"product_id" gorm:"primaryKey;autoIncrement"`
	UserID        uint           `json:"user_id"`
	CategoryID    uint           `json:"category_id"`
	ProductName   string         `json:"product_name"`
	Description   string         `json:"description"`
	Price         float64        `json:"price"`
	Quantity      uint           `json:"quantity"`
	ProductImages []ProductImage `json:"product_images" gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ProductImageID uint `json:"product_image_id" gorm:"primaryKey;autoIncrement"`
	ProductID      uint `json:"product_id"`
	ImageID        uint `json:"image_id" gorm:"foreignKey:ImageID"`
	Image          Image
}

type Image struct {
	ImageID   uint   `json:"image_id" gorm:"primaryKey;autoIncrement"`
	BucketKey string `json:"bucket_key"`
	Path      string `json:"path"`
}

func (DetailProductRes) TableName() string {
	return "products"
}

func (ProductImage) TableName() string {
	return "product_images"
}
func (Image) TableName() string {
	return "images"
}
