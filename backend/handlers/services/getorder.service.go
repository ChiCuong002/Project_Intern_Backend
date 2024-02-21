package services

import (
	"errors"
	"main/schema"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PurchaseProduct(db *gorm.DB, c echo.Context, userID, productID, quantity uint) error {
	var user schema.User
	var product schema.Product
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := db.First(&product, productID).Error; err != nil {
		return err
	}

	// Kiểm tra số dư người dùng và sự có sẵn của sản phẩm
	if user.Balance < product.Price*float64(quantity) || product.Quantity < quantity {
		return errors.New("insufficient balance or insufficient product quantity")
	}

	// Thực hiện giao dịch mua hàng
	orderTotal := product.Price * float64(quantity)
	user.Balance -= orderTotal
	product.Quantity -= quantity

	// Cập nhật số dư trong bảng User
	if err := db.Model(&user).Update("balance", user.Balance).Error; err != nil {
		return err
	}

	// Thêm mới đơn hàng vào bảng Order
	newOrder := schema.Order{
		UserID:      userID,
		ProductID:   productID,
		OrderTotal:  orderTotal,
		OrderDate:   time.Now(),
		OrderStatus: "Đã đặt hàng",
	}

	if err := db.Create(&newOrder).Error; err != nil {
		return err
	}
	// Không cần trả về số dư cập nhật, chỉ trả về nil nếu không có lỗi
	return c.JSON(http.StatusOK, echo.Map{
		"user":    user,
		"message": "Dat hang thành công",
	})
}
