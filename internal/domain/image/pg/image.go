package pg

import (
	"context"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/errorx"
	"go.uber.org/zap"
)

// ImageRepository
type ImageRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewImageRepository
func NewImageRepository(client *ent.Client, logger *zap.Logger) *ImageRepository {
	return &ImageRepository{client: client, logger: logger}
}

// BulkCreateTx
func (r *ImageRepository) BulkCreateTx(
	ctx context.Context,
	images image.Images,
) (image.Images, error) {
	if len(images) == 0 {
		return nil, nil
	}
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		r.logger.Error("using tx in non tx context", zap.Error(errorx.ErrInternal))
		return nil, errorx.ErrInternal
	}
	client := tx.Client()
	rows, err := client.Image.
		CreateBulk(mapImagesToCreateBuilders(client, images)...).
		Save(ctx)
	if err != nil {
		r.logger.Error("bulk create images error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainImagesFromRows(rows), nil
}

func mapImagesToCreateBuilders(
	client *ent.Client,
	images image.Images,
) []*ent.ImageCreate {
	builders := make([]*ent.ImageCreate, 0, len(images))
	for _, image := range images {
		builder := client.Image.Create().
			SetName(image.Props.Name).
			SetOriginalName(image.Props.OriginalName)
		builders = append(builders, builder)
	}
	return builders
}

func mapDomainImageFromRow(row *ent.Image) *image.Image {
	return &image.Image{
		ID:        image.ID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &image.Props{
			Name:         row.Name,
			OriginalName: row.OriginalName,
		},
	}
}

func mapDomainImagesFromRows(rows ent.Images) image.Images {
	result := make(image.Images, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainImageFromRow(row))
	}
	return result
}
