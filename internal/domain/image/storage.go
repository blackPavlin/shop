package image

import (
	"context"
	"io"
)

//go:generate mockgen -source $GOFILE -destination "storage_mock.go" -package "image"

// Storage represents file store.
type Storage interface {
	Upload(ctx context.Context, props *StorageProps) error
	BulkUpload(ctx context.Context, props []*StorageProps) error
	Remove(ctx context.Context, name string) error
	BulkRemove(ctx context.Context, names []string) error
}

// StorageProps represents image storage fields.
type StorageProps struct {
	Name        string
	Reader      io.Reader
	Size        int64
	ContentType string
}
