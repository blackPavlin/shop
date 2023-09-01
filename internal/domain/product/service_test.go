package product_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"

	"github.com/blackPavlin/shop/internal/domain/image"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/repositoryx"
	"github.com/blackPavlin/shop/pkg/testutilx"
)

func TestCartServiceSuite(t *testing.T) {
	suite.Run(t, new(ProductServiceSuite))
}

type ProductServiceBaseSuite struct {
	testutilx.BaseSuite
	productRepo      *product.MockRepository
	productImageRepo *product.MockImageRepository
	imageStorage     *image.MockStorage
	txManages        repositoryx.TxManager
}

func (s *ProductServiceBaseSuite) SetupTest() {
	s.BaseSuite.SetupTest()
	s.productRepo = product.NewMockRepository(s.Ctrl)
	s.productImageRepo = product.NewMockImageRepository(s.Ctrl)
	s.imageStorage = image.NewMockStorage(s.Ctrl)
	s.txManages = repositoryx.NewNopTxManager()
}

type ProductServiceSuite struct {
	ProductServiceBaseSuite
	productService product.Service
}

func (s *ProductServiceSuite) SetupTest() {
	s.ProductServiceBaseSuite.SetupTest()
	s.productService = product.NewUseCase(
		s.productRepo,
		s.productImageRepo,
		s.imageStorage,
		s.txManages,
	)
}

func (s *ProductServiceSuite) DeleteTest() {
	type args struct {
		ctx       context.Context
		productID product.ID
	}
	type want struct {
		err error
	}
	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			args.ctx = prepare(args.ctx)

			err := s.productService.Delete(args.ctx, args.productID)
			s.ErrorIs(err, want.err)
		}
	}
	tests := []struct {
		name    string
		prepare func(ctx context.Context) context.Context
		args    args
		want    want
	}{
		{
			name: "success",
			prepare: func(ctx context.Context) context.Context {
				s.productImageRepo.EXPECT().
					Query(ctx, &product.ImageQueryCriteria{
						Filter: product.ImageFilter{
							ProductID: product.IDFilter{Eq: product.IDs{1}},
						},
					}).
					Return(product.Images{
						{
							ID: 1,
							Props: &product.ImageProps{
								Name: "74b9a255-af1f-4222-910d-7f477b559608.png",
							},
						},
						{
							ID: 2,
							Props: &product.ImageProps{
								Name: "a592cbe8-018c-4595-a2e0-a80dd76189d0.png",
							},
						},
					}, nil)
				s.productRepo.EXPECT().
					Delete(ctx, &product.Filter{
						ID: product.IDFilter{Eq: product.IDs{1}},
					}).
					Return(nil)
				s.imageStorage.EXPECT().
					BulkRemove(ctx, []string{
						"74b9a255-af1f-4222-910d-7f477b559608.png",
						"a592cbe8-018c-4595-a2e0-a80dd76189d0.png",
					}).
					Return(nil)
				return ctx
			},
			args: args{
				ctx:       context.Background(),
				productID: 1,
			},
			want: want{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}
