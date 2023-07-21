package image

import (
	"context"
	"database/sql"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "image"

// Service represents image use cases.
type Service interface {
	BulkCreate(ctx context.Context, props []*StorageProps) (Images, error)
}

// UseCase represents image service.
type UseCase struct {
	logger *zap.Logger

	imageRepo    Repository
	imageStorage Storage
	txManager    repositoryx.TxManager
}

// NewUseCase create instance of UseCase.
func NewUseCase(
	logger *zap.Logger,
	imageRepo Repository,
	imageStorage Storage,
	txManager repositoryx.TxManager,
) *UseCase {
	return &UseCase{
		logger:       logger,
		imageRepo:    imageRepo,
		imageStorage: imageStorage,
		txManager:    txManager,
	}
}

// BulkCreate images.
func (s *UseCase) BulkCreate(ctx context.Context, props []*StorageProps) (Images, error) {
	if len(props) == 0 {
		return nil, nil
	}
	var (
		images = make(Images, 0, len(props))
		err    error
	)
	for _, prop := range props {
		id, err := uuid.NewUUID()
		if err != nil {
			s.logger.Error("generate uuid error", zap.Error(err))
			return nil, errorx.ErrInternal
		}
		extension := filepath.Ext(prop.Name)
		name := fmt.Sprintf("%s%s", id.String(), extension)
		images = append(images, &Image{
			Props: &Props{
				OriginalName: prop.Name,
				Name:         name,
			},
		})
		prop.Name = name
		prop.ContentType = mime.TypeByExtension(extension)
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
