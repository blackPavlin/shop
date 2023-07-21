// Package storage contains implementations for file storage
package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/s3x"
)

// Storage s3 storage implementation.
type Storage struct {
	client *minio.Client
	config *s3x.S3Config
	logger *zap.Logger
}

// NewStorage create instance of Storage.
func NewStorage(client *minio.Client, config *s3x.S3Config, logger *zap.Logger) *Storage {
	return &Storage{client: client, config: config, logger: logger}
}

// Upload file to storage.
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
