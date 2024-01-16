package schema

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User `gorm:"foreignKey:RoleID"`
}
type User struct {
	UserID      uint `gorm:"primaryKey;autoIncrement"`
	RoleID      uint
	FirstName   string
	LastName    string
	Address     string
	Email       string
	PhoneNumber string
	Password    string
	Products    []Product `gorm:"foreignKey:UserID"`
	Orders      []Order   `gorm:"foreignKey:UserID"`
}
type Category struct {
	CategoryID   uint `gorm:"primaryKey;autoIncrement"`
	CategoryName string
	Products     []Product `gorm:"foreignKey:CategoryID"`
}
type Product struct {
	ProductID   uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	CategoryID  uint
	ProductName string
	Description string
	Price       float64
	Quantity    uint
	ImagePath   string
	Images      []Image `gorm:"foreignKey:ProductID"`
	Orders      []Order `gorm:"foreignKey:ProductID"`
}
type Image struct {
	ImageID   uint `gorm:"primaryKey;autoIncrement"`
	ProductID uint
	Path      string
}
type Order struct {
	OrderID     uint `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	ProductID   uint
	OrderTotal  float64
	OrderDate   time.Time `gorm:"type:timestamp"`
	OrderStatus string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (u *User) ChangePassword(db *gorm.DB, newPassword string) error {
	hash, err := HashPassword(newPassword)
	if err != nil {
		panic("ma hoa that bai")
	}
	match := CheckPasswordHash(newPassword, hash)
	if !match {
		panic("mk ma hoa khong trung")
	}
	// Cập nhật mật khẩu người dùng trong cơ sở dữ liệu
	result := db.Model(u).Update("password", hash)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func Migration() {
	dsn := "host=localhost user=postgres password=chicuong dbname=fitness-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Role{}, &User{}, &Category{}, &Product{}, &Image{}, &Order{})
}
