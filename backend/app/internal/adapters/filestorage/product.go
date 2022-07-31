package filestorage

import "github.com/minio/minio-go/v7"

type ProductFileStorage struct {
	client *minio.Client
}

func NewProductFileStorage(client *minio.Client) *ProductFileStorage {
	return &ProductFileStorage{
		client: client,
	}
}
