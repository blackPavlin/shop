// Package storage contains implementations for file storage
package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

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

// BulkUpload files to storage.
func (s *Storage) BulkUpload(ctx context.Context, props []*image.StorageProps) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, prop := range props {
		prop := prop
		g.Go(func() error {
			if err := s.Upload(ctx, prop); err != nil {
				s.logger.Error("bulkUpload image to storage error", zap.Error(err))
				return errorx.ErrInternal
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return errorx.ErrInternal
	}
	return nil
}

// Remove file from storage.
func (s *Storage) Remove(ctx context.Context, name string) error {
	err := s.client.RemoveObject(ctx, s.config.BucketName, name, minio.RemoveObjectOptions{})
	if err != nil {
		s.logger.Error("remove image from storage error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
}

// BulkRemove files from storage.
func (s *Storage) BulkRemove(ctx context.Context, names []string) error {
	g, ctx := errgroup.WithContext(ctx)
	for _, name := range names {
		name := name
		g.Go(func() error {
			if err := s.Remove(ctx, name); err != nil {
				s.logger.Error("bulkRemove image from storage error", zap.Error(err))
				return errorx.ErrInternal
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return errorx.ErrInternal
	}
	return nil
}
