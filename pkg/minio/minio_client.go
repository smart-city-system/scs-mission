package minio_client

import (
	"context"
	"mime/multipart"
	"scs-guard/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	client     *minio.Client
	logger     logger.Logger
	bucketName string
}

func NewMinioClient(endpoint string, accessKeyID string, secretAccessKey string, bucketName string, logger logger.Logger) (*MinioClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	if err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{}); err != nil {
		exists, errBucketExists := client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			return &MinioClient{client: client, logger: logger}, nil
		} else {
			return nil, err
		}
	}

	return &MinioClient{client: client, logger: logger}, nil
}

func (c *MinioClient) GetClient() *minio.Client {
	return c.client
}

func (c *MinioClient) UploadFile(bucketName string, objectName string, file multipart.File, fileSize int64) (minio.UploadInfo, error) {
	info, err := c.client.PutObject(context.Background(), bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		c.logger.Errorf("Failed to upload file to Minio: %v", err)
		return minio.UploadInfo{}, err
	}
	return info, nil
}
