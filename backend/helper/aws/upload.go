package aws

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
	if contentType != "image/jpeg" && contentType != "image/png" {
		return fmt.Errorf("Invalid file type. Only images with JPG,JPEG,JPG format are allowed !")
	}
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
func GetImagePath(bucketKey string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", os.Getenv("BUCKETNAME"), os.Getenv("REGION"), bucketKey)
}
func ConfigureAWSService() (awsService, error) {
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
func (awsSvc awsService) UpdateFile(bucketName, bucketKey string, file multipart.File) error {
	// Read the first 512 bytes of the file
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		log.Println("Error while reading the file ", err)
		return err
	}
	// Detect the Content-Type of the file
	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" {
		return fmt.Errorf("Invalid file type. Only images with JPG,JPEG,JPG format are allowed !")
	}
	// Reset the read pointer to the start of the file
	file.Seek(0, 0)

	// Upload the new file to S3 with the same key
	_, err = awsSvc.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(bucketKey),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Println("Error while updating the file ", err)
	}
	return err
}
