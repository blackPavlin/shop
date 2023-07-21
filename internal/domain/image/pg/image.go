// Package pg contains implementations for image repositories
package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entimage "github.com/blackPavlin/shop/ent/image"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/errorx"
)

// ImageRepository pg repository implementation.
type ImageRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewImageRepository create instance of ImageRepository.
func NewImageRepository(client *ent.Client, logger *zap.Logger) *ImageRepository {
	return &ImageRepository{client: client, logger: logger}
}

// Get image from db.
func (r *ImageRepository) Get(ctx context.Context, filter *image.Filter) (*image.Image, error) {
	row, err := r.client.Image.Query().
		Where(makePredicates(filter)...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get image error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainImageFromRow(row), nil
}

// BulkCreateTx images in db.
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

func makePredicates(filter *image.Filter) []predicate.Image {
	predicates := make([]predicate.Image, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entimage.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entimage.IDNotIn(filter.ID.Neq.ToInt64()...))
	}
	return predicates
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
