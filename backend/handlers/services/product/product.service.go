package service

import (
	//"fmt"
	//storage "main/database"

	"fmt"
	storage "main/database"
	"main/helper/scope"
	paginationHelper "main/helper/struct"
	helper "main/helper/struct/product"
	productHelper "main/helper/struct/product"
	"main/schema"

	"gorm.io/gorm"
)

const (
	STATUS_ACTIVE   uint = 1
	STATUS_INACTIVE uint = 2
)

func InsertProduct(tx *gorm.DB, product *helper.ProductInsert) error {
	product.StatusID = STATUS_ACTIVE
	result := tx.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func InsertImage(tx *gorm.DB, image *helper.ImageInsert) error {
	result := tx.Create(image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func InsertProductImage(tx *gorm.DB, productImg *helper.ProductImageInsert) error {
	result := tx.Create(productImg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateProduct(product *helper.UpdateProduct) error {
	db := storage.GetDB()
	result := db.First(&schema.Product{}, product.ProductID)
	if result.Error != nil {
		return fmt.Errorf("invalid product id")
	}
	result = db.First(&schema.Category{}, product.CategoryID)
	if result.Error != nil {
		return fmt.Errorf("invalid category id")
	}
	result = db.Model(&schema.Product{}).Where("product_id = ?", product.ProductID).Updates(map[string]interface{}{
		"category_id":  product.CategoryID,
		"product_name": product.ProductName,
		"description":  product.Description,
		"price":        product.Price,
		"quantity":     product.Quantity})
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())
	}
	return nil
}
func DetailProduct(id uint) (productHelper.DetailProductRes, error) {
	db := storage.GetDB()
	product := productHelper.DetailProductRes{}
	err := db.Model(&productHelper.DetailProductRes{}).Preload("ProductImages.Image").Preload("User").Preload("Category").Preload("Status").First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
func SearchProducts(query *gorm.DB, search string) *gorm.DB {
	if search != "" {
		query = query.Where("product_name like ?", "%"+search+"%")
	}
	return query

}
func GetAllProduct(pagination paginationHelper.Pagination) (*paginationHelper.Pagination, error) {
	db := storage.GetDB()
	products := []productHelper.DetailProductRes{}
	query := db.Model(&products)
	query = SearchProducts(query, pagination.Search)
	query = query.Scopes(scope.Paginate(query, &pagination))
	query.Preload("ProductImages.Image").Preload("User").Preload("Category").Preload("Status").Find(&products)
	pagination.Rows = products
	return &pagination, nil
}
func GetMyInventory(pagination paginationHelper.Pagination, userID uint) (*paginationHelper.Pagination, error) {
	db := storage.GetDB()
	products := []productHelper.DetailProductRes{}
	query := db.Model(&products).Where("user_id = ? and status_id = 2", userID)
	fmt.Printf("query: %+v\n", query)
	query = SearchProducts(query, pagination.Search)
	query = query.Scopes(scope.Paginate(query, &pagination))
	query.Preload("ProductImages.Image").Preload("User").Preload("Category").Preload("Status").Find(&products)
	pagination.Rows = products
	return &pagination, nil
}
func GetMyProduct(pagination paginationHelper.Pagination, userID uint) (*paginationHelper.Pagination, error) {
	db := storage.GetDB()
	products := []productHelper.DetailProductRes{}
	query := db.Model(&products).Where("user_id = ?", userID).Where("status_id = 1")
	fmt.Printf("query: %+v\n", query)
	query = SearchProducts(query, pagination.Search)
	query = query.Scopes(scope.Paginate(query, &pagination))
	query.Preload("ProductImages.Image").Preload("User").Preload("Category").Preload("Status").Where(&schema.Product{UserID: userID, StatusID: STATUS_ACTIVE}).Find(&products)
	pagination.Rows = products
	return &pagination, nil
}
func CompareUserID(userID, productID uint) error {
	db := storage.GetDB()
	product := schema.Product{}
	err := db.First(&product, productID).Error
	if err != nil {
		return fmt.Errorf("can't find the product. Check product id again")
	}
	if userID != product.UserID {
		return fmt.Errorf("user is not own this product")
	}
	return nil
}

func BlockProduct(id uint) error {
	db := storage.GetDB()
	product := helper.DetailProductRes{}
	result := db.Model(&product).First(&product, "product_id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("falied to get product")
	}
	if product.StatusID != STATUS_ACTIVE {
		return fmt.Errorf("this product is currently inactive")
	}
	product.StatusID = STATUS_INACTIVE
	result = db.Model(&product).Where("product_id", product.ProductID).Update("status_id", product.StatusID)
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())
	}
	return nil
}

func UnblockProduct(id uint) error {
	db := storage.GetDB()
	product := helper.DetailProductRes{}
	result := db.Model(&product).First(&product, "product_id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("falied to get product. %s", result.Error.Error())
	}
	if product.StatusID != STATUS_INACTIVE {
		return fmt.Errorf("this product is currently active")
	}
	product.StatusID = STATUS_ACTIVE
	result = db.Model(&product).Where("product_id", product.ProductID).Update("status_id", product.StatusID)
	if result.Error != nil {
		return fmt.Errorf(result.Error.Error())
	}
	return nil
}
func MyProduct() error {
	return nil
}
