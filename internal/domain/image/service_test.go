package image_test

import (
	"context"
	"testing"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/pkg/repositoryx"
	"github.com/blackPavlin/shop/pkg/testutilx"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestImageServiceSuite(t *testing.T) {
	suite.Run(t, new(ImageServiceSuite))
}

type ImageServiceBaseSuite struct {
	testutilx.BaseSuite
	logger       *zap.Logger
	imageRepo    *image.MockRepository
	imageStorage *image.MockStorage
	txManager    repositoryx.TxManager
}

func (s *ImageServiceBaseSuite) SetupTest() {
	s.BaseSuite.SetupTest()
	s.logger = zaptest.NewLogger(s.T())
	s.imageRepo = image.NewMockRepository(s.Ctrl)
	s.imageStorage = image.NewMockStorage(s.Ctrl)
	s.txManager = repositoryx.NopTxManager{}
}

type ImageServiceSuite struct {
	ImageServiceBaseSuite
	ImageService image.Service
}

func (s *ImageServiceSuite) SetupTest() {
	s.ImageServiceBaseSuite.SetupTest()
	s.ImageService = image.NewUseCase(s.logger, s.imageRepo, s.imageStorage, s.txManager)
}

func (s *ImageServiceSuite) TestBulkCreate() {
	type args struct {
		ctx   context.Context
		props []*image.StorageProps
	}
	type want struct {
		res image.Images
		err error
	}
	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {}
	}
	tests := []struct {
		name    string
		prepare func(ctx context.Context) context.Context
		args    args
		want    want
	}{}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}
