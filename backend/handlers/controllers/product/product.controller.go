package controllers

import (
	"context"
	"fmt"
	storage "main/database"
	service "main/handlers/services/product"
	helper "main/helper/struct/product"
	"strconv"

	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type productForm struct {
	CategoryID  uint                    `json:"category_id" form:"category_id"`
	ProductName string                  `json:"product_name" form:"product_name"`
	Description string                  `json:"description" form:"description"`
	Price       float64                 `json:"price" form:"price"`
	Quantity    uint                    `json:"quantity" form:"quantity"`
	Images      []*multipart.FileHeader `json:"images" form:"images"`
}
type awsService struct {
	S3Client *s3.Client
}

func (awsSvc awsService) UploadFile(bucketName, bucketKey string, file multipart.File) error {
	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		log.Println("Error while reading the file ", err)
		return err
	}
	// Detect the Content-Type of the file
	contentType := http.DetectContentType(buffer)

	// Reset the read pointer to the start of the file
	file.Seek(0, 0)

	_, err = awsSvc.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(bucketKey),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Println("Error while uploading the file ", err)
	}
	return err
}
func getImagePath(bucketKey string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", os.Getenv("BUCKETNAME"), os.Getenv("REGION"), bucketKey)
}
func configureAWSService() (awsService, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("REGION")),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)
	if err != nil {
		return awsService{}, err
	}
	return awsService{S3Client: s3.NewFromConfig(cfg)}, nil
}
func AddProduct(c echo.Context) error {
	productData := &productForm{}
	if err := c.Bind(productData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request data")
	}
	//begin transaction
	tx := storage.GetDB().Begin()
	//create product type
	product := &helper.ProductInsert{
		UserID:      c.Get("userID").(uint),
		CategoryID:  productData.CategoryID,
		ProductName: productData.ProductName,
		Description: productData.Description,
		Price:       productData.Price,
		Quantity:    productData.Quantity,
	}
	// //insert product into database
	err := service.InsertProduct(tx, product)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Failed to insert product")
	}
	// //Get multipart form from the request
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid form data")
	}
	// Get the image files from the form
	images, ok := form.File["images"]
	if !ok {
		return c.JSON(http.StatusBadRequest, "No image files")
	}
	//load .env variable
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Configure the AWS service
	awsService, err := configureAWSService()
	if err != nil {
		log.Println("Error while configuring the AWS service:", err)
	}
	// Iterate over the image files and upload each one
	for _, image := range images {
		// Generate a unique key for the S3 bucket
		bucketKey := uuid.New().String()
		//create image path
		imagePath := getImagePath(bucketKey)
		//insert the image into database
		img := &helper.ImageInsert{
			BucketKey: bucketKey,
			Path:      imagePath,
		}
		err = service.InsertImage(tx, img)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Failed to insert image")
		}
		productImg := &helper.ProductImageInsert{
			ProductID: product.ProductID,
			ImageID:   img.ImageID,
		}
		err = service.InsertProductImage(tx, productImg)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Failed to insert product image")
		}
		//product.Images = append(product.Images, *img)
		// Open the image file
		src, err := image.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer src.Close()
		// Upload the image to AWS
		err = awsService.UploadFile(os.Getenv("BUCKETNAME"), bucketKey, src)
		if err != nil {
			log.Println("Failed to upload the image:", err)
		} else {
			log.Println("Image uploaded successfully")
		}
	}
	tx.Commit()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Add product successfully",
		"product": product,
	})
}
func DetailProduct(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failed to get product id",
		})
	}
	product, err := service.DetailProduct(uint(productID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"product": product,
	})
}
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	product := helper.UpdateProduct{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	product.ProductID = uint(id)
	err = service.UpdateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Update product data successfully",
	})
}
