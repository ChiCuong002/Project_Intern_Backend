package helper

import "mime/multipart"

type UpdateData struct {
	UserID    uint
	FirstName string                `form:"first_name"`
	LastName  string                `form:"last_name"`
	Address   string                `form:"address"`
	Email     string                `form:"email"`
	Image     *multipart.FileHeader `form:"image"`
}
type UserInsert struct {
	UserID    uint
	FirstName string
	LastName  string
	Address   string
	Email     string
	Image     uint
}

func (UserInsert) TableName() string {
	return "users"
}

type UserResponse struct {
	UserID      uint    `json:"user_id"`
	RoleID      uint    `json:"role_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Address     string  `json:"address"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Balance     float64 `json:"balance"`
	IsActive    bool    `json:"is_active"`
	ImageID     uint    `json:"-" gorm:"foreignKey:ImageID"`
	Image       *Image  `json:"image"`
}
type Image struct {
	ImageID   uint   `json:"image_id" gorm:"primaryKey;autoIncrement"`
	BucketKey string `json:"bucket_key"`
	Path      string `json:"image_path"`
}

func (UserResponse) TableName() string {
	return "users"
}

type AddBalanceReq struct {
	UserID      uint    `json:"user_id"`
	AmountToAdd float64 `json:"amount"`
}
type Gif struct {
	Amount float64 `json:"amount"`
}
