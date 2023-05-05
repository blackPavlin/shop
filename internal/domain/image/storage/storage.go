package storage

import (
	"context"

	"github.com/minio/minio-go/v7"
)

// Storage
type Storage struct {
	client *minio.Client
}

// NewStorage
func NewStorage(client *minio.Client) *Storage {
	return &Storage{client: client}
}

// Upload
func (s *Storage) Upload(ctx context.Context) error {
	return nil
}
