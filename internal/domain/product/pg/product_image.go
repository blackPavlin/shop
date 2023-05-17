package pg

import (
	"context"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	entproductimage "github.com/blackPavlin/shop/ent/productimage"
	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/domain/product"
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
	images product.Images,
) (product.Images, error) {
	if len(images) == 0 {
		return nil, nil
	}
	tx := ent.TxFromContext(ctx)
	if tx == nil {
		r.logger.Error("using tx in non tx context", zap.Error(errorx.ErrInternal))
		return nil, errorx.ErrInternal
	}
	client := tx.Client()
	rows, err := client.ProductImage.
		CreateBulk(mapImagesToCreateBuilders(client, images)...).
		Save(ctx)
	if err != nil {
		// todo: IsForeignKeyViolationErr
		r.logger.Error("bulk create product images error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainImagesFromRows(rows), nil
}

// Query
func (r *ImageRepository) Query(
	ctx context.Context,
	criteria *product.ImageQueryCriteria,
) (product.Images, error) {
	rows, err := r.client.ProductImage.
		Query().
		Where(makeImagePredicates(&criteria.Filter)...).
		All(ctx)
	if err != nil {
		r.logger.Error("query product images error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainImagesFromRows(rows), nil
}

func mapImagesToCreateBuilders(
	client *ent.Client,
	images product.Images,
) []*ent.ProductImageCreate {
	builders := make([]*ent.ProductImageCreate, 0)
	for _, image := range images {
		builder := client.ProductImage.
			Create().
			SetProductID(int64(image.Props.ProductID))
		builders = append(builders, builder)
	}
	return builders
}

func makeImagePredicates(filter *product.ImageFilter) []predicate.ProductImage {
	predicates := make([]predicate.ProductImage, 0)
	if len(filter.ProductID.Eq) > 0 {
		predicates = append(
			predicates,
			entproductimage.ProductIDIn(filter.ProductID.Eq.ToInt64()...),
		)
	}
	if len(filter.ProductID.Neq) > 0 {
		predicates = append(
			predicates,
			entproductimage.ProductIDNotIn(filter.ProductID.Neq.ToInt64()...),
		)
	}
	return predicates
}

func mapDomainImageFromRow(row *ent.ProductImage) *product.Image {
	return &product.Image{
		ID:        product.ImageID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &product.ImageProps{
			ProductID: product.ID(row.ProductID),
			ImageID:   image.ID(row.ImageID),
		},
	}
}

func mapDomainImagesFromRows(rows ent.ProductImages) product.Images {
	result := make(product.Images, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainImageFromRow(row))
	}
	return result
}
