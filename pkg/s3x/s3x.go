package s3x

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// S3Config
type S3Config struct {
	Endpoint        string `envconfig:"ENDPOINT" required:"true"`
	AccessKeyID     string `envconfig:"ACCESS_KEY_ID" required:"true"`
	SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY" required:"true"`
	UseSSL          bool   `envconfig:"USE_SSL" required:"true"`
	BucketName      string `envconfig:"BUCKET_NAME" required:"true"`
	Location        string `envconfig:"LOCATION" required:"true"`
}

// NewStorage
func NewStorage(ctx context.Context, config *S3Config) (*minio.Client, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}
	if err := client.MakeBucket(ctx, config.BucketName, minio.MakeBucketOptions{
		Region: config.Location,
	}); err != nil {
		exists, err := client.BucketExists(ctx, config.BucketName)
		if err == nil && exists {
			return client, nil
		}
		return nil, err
	}
	return client, nil
}
