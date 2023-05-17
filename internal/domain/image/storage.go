package image

import "context"

//go:generate mockgen -source $GOFILE -destination "storage_mock.go" -package "image"

// Storage
type Storage interface {
	Upload(ctx context.Context, props *StorageProps) error
}

type StorageProps struct {
	Name        string
	Buffer      []byte
	ContentType string
}
