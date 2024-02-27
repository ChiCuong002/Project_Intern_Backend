package services

import (
	"errors"
	"main/schema"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PurchaseProduct(db *gorm.DB, c echo.Context, userID, productID uint) error {
	var user schema.User
	var product schema.Product
	var seller schema.User

	if userID == productID {
		return errors.New("Bạn không thể mua sản phẩm của chính bạn")
	}

	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := db.First(&product, productID).Error; err != nil {
		return err
	}

	// Kiểm tra số dư người mua và sự có sẵn của sản phẩm
	if user.Balance < product.Price {
		return errors.New("Số dư không đủ hoặc số lượng sản phẩm không đủ")
	}

	// Kiểm tra người bán và xác nhận không mua sản phẩm của chính mình
	if err := db.First(&seller, product.UserID).Error; err != nil {
		return errors.New("Người bán không tồn tại")
	}

	if userID == seller.UserID {
		return errors.New("Bạn không thể mua sản phẩm của chính bạn")
	}

	// Thực hiện giao dịch mua hàng
	orderTotal := product.Price
	user.Balance -= orderTotal
	seller.Balance += orderTotal

	// Cập nhật số dư trong bảng User
	if err := db.Model(&user).Update("balance", user.Balance).Error; err != nil {
		return err
	}

	if err := db.Model(&seller).Update("balance", seller.Balance).Error; err != nil {
		return err
	}

	// Set status_id của sản phẩm thành 2
	product.StatusID = 2

	if err := db.Model(&product).Updates(map[string]interface{}{"status_id": product.StatusID}).Error; err != nil {
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

	// Trả về thông tin người mua và thông báo thành công
	return c.JSON(http.StatusOK, echo.Map{
		"user":    user,
		"message": "Đặt hàng thành công",
	})
}
