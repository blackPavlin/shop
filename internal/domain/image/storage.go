package image

import "context"

//go:generate mockgen -source $GOFILE -destination "storage_mock.go" -package "image"

// Storage represents file store.
type Storage interface {
	Upload(ctx context.Context, props *StorageProps) error
	Remove(ctx context.Context, name string) error
}

type StorageProps struct {
	Name        string
	Buffer      []byte
	ContentType string
}
