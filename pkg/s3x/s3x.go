// Package s3x contains utils to work with S3 storages.
package s3x

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// S3Config is s3 storage configuration.
type S3Config struct {
	Endpoint        string `envconfig:"ENDPOINT" required:"true"`
	AccessKeyID     string `envconfig:"ACCESS_KEY_ID" required:"true"`
	SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY" required:"true"`
	UseSSL          bool   `envconfig:"USE_SSL" required:"true"`
	BucketName      string `envconfig:"BUCKET_NAME" required:"true"`
	Location        string `envconfig:"LOCATION" required:"true"`
}

// NewStorage create instance of s3 storage.
func NewStorage(ctx context.Context, config *S3Config) (*minio.Client, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("connect to s3 bucket error: %w", err)
	}
	if err := client.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{
		Region: config.Location,
	}); err != nil {
		exists, err := client.BucketExists(ctx, config.BucketName)
		if err == nil && exists {
			return client, nil
		}
		return nil, fmt.Errorf("check exists bucket error: %w", err)
	}
	return client, nil
}
