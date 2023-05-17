package storage

import (
	"bytes"
	"context"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/s3x"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

// Storage
type Storage struct {
	client *minio.Client
	config *s3x.S3Config
	logger *zap.Logger
}

// NewStorage
func NewStorage(client *minio.Client, config *s3x.S3Config, logger *zap.Logger) *Storage {
	return &Storage{client: client, config: config, logger: logger}
}

// Upload
func (s *Storage) Upload(ctx context.Context, props *image.StorageProps) error {
	_, err := s.client.PutObject(
		ctx,
		s.config.BucketName,
		props.Name,
		bytes.NewReader(props.Buffer),
		int64(len(props.Buffer)),
		minio.PutObjectOptions{
			ContentType: props.ContentType,
		},
	)
	if err != nil {
		s.logger.Error("upload image to storage error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
}
