package utils

import (
    "context"
    "fmt"
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
    "video-platform-backend/config"
)

var MinioClient *minio.Client

func InitMinio() {
    var err error
    MinioClient, err = minio.New(config.AppConfig.Minio.Endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(config.AppConfig.Minio.AccessKey, config.AppConfig.Minio.SecretKey, ""),
        Secure: config.AppConfig.Minio.UseSSL,
    })
    if err != nil {
        log.Fatalf("Failed to initialize MinIO client: %v", err)
    }

    // 确保 bucket 存在
    ctx := context.Background()
    exists, err := MinioClient.BucketExists(ctx, config.AppConfig.Minio.Bucket)
    if err != nil {
        log.Fatalf("Failed to check bucket: %v", err)
    }
    if !exists {
        err = MinioClient.MakeBucket(ctx, config.AppConfig.Minio.Bucket, minio.MakeBucketOptions{})
        if err != nil {
            log.Fatalf("Failed to create bucket: %v", err)
        }
        fmt.Println("Created bucket:", config.AppConfig.Minio.Bucket)
    }
}
