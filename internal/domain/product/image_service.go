package product

import (
	"context"
	"database/sql"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate mockgen -source $GOFILE -destination "image_service_mock.go" -package "product"

// ImageService represents product image use cases.
type ImageService interface {
	BulkCreate(ctx context.Context, productID ID, props []*image.StorageProps) (Images, error)
	Delete(ctx context.Context, imageID ImageID) error
}

// ImageUseCase represents product image service.
type ImageUseCase struct {
	logger       *zap.Logger
	productRepo  Repository
	imageRepo    ImageRepository
	imageStorage image.Storage
	txManager    repositoryx.TxManager
}

// NewImageUseCase create instance of ImageUseCase.
func NewImageUseCase(
	logger *zap.Logger,
	productRepo Repository,
	imageRepo ImageRepository,
	imageStorage image.Storage,
	txManager repositoryx.TxManager,
) *ImageUseCase {
	return &ImageUseCase{logger, productRepo, imageRepo, imageStorage, txManager}
}

// BulkCreate product images.
func (s *ImageUseCase) BulkCreate(
	ctx context.Context,
	productID ID,
	props []*image.StorageProps,
) (Images, error) {
	if len(props) == 0 {
		return nil, nil
	}
	var (
		images = make(Images, 0, len(props))
		err    error
	)
	_, err = s.productRepo.Get(ctx, &Filter{
		ID: IDFilter{Eq: IDs{productID}},
	})
	if err != nil {
		return nil, fmt.Errorf("get product error: %w", err)
	}
	for _, prop := range props {
		id, err := uuid.NewUUID()
		if err != nil {
			s.logger.Error("generate uuid error", zap.Error(err))
			return nil, errorx.ErrInternal
		}
		extension := filepath.Ext(prop.Name)
		prop.Name = fmt.Sprintf("%s%s", id.String(), extension)
		prop.ContentType = mime.TypeByExtension(extension)
		images = append(images, &Image{
			Props: &ImageProps{productID, prop.Name},
		})
	}
	err = s.txManager.RunTransaction(ctx, &sql.TxOptions{}, func(ctx context.Context) error {
		images, err = s.imageRepo.BulkCreateTx(ctx, images)
		if err != nil {
			return fmt.Errorf("bulkCreate images error: %w", err)
		}
		for _, prop := range props {
			if err := s.imageStorage.Upload(ctx, prop); err != nil {
				return fmt.Errorf("upload images error: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("bulkCreate images transaction error: %w", err)
	}
	return images, nil
}

// Delete image from db and file storage.
func (s *ImageUseCase) Delete(ctx context.Context, imageID ImageID) error {
	img, err := s.imageRepo.Get(ctx, &ImageFilter{
		ImageID: ImageIDFilter{Eq: ImageIDs{imageID}},
	})
	if err != nil {
		return fmt.Errorf("get image error: %w", err)
	}
	err = s.txManager.RunTransaction(ctx, &sql.TxOptions{}, func(ctx context.Context) error {
		if err := s.imageRepo.DeleteTx(ctx, imageID); err != nil {
			return fmt.Errorf("delete images error: %w", err)
		}
		if err := s.imageStorage.Remove(ctx, img.Props.Name); err != nil {
			return fmt.Errorf("remove images from storage error: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("delete images transaction error: %w", err)
	}
	return nil
}
