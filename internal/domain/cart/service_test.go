package cart_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/testutilx"
)

func TestCartServiceSuite(t *testing.T) {
	suite.Run(t, new(CartServiceSuite))
}

type CartServiceBaseSuite struct {
	testutilx.BaseSuite
	cartRepo    *cart.MockRepository
	productRepo *product.MockRepository
}

func (s *CartServiceBaseSuite) SetupTest() {
	s.BaseSuite.SetupTest()
	s.cartRepo = cart.NewMockRepository(s.Ctrl)
	s.productRepo = product.NewMockRepository(s.Ctrl)
}

type CartServiceSuite struct {
	CartServiceBaseSuite
	cartService cart.Service
}

func (s *CartServiceSuite) SetupTest() {
	s.CartServiceBaseSuite.SetupTest()
	s.cartService = cart.NewUseCase(s.cartRepo, s.productRepo)
}

func (s *CartServiceSuite) TestSave() {
	type args struct {
		ctx   context.Context
		props *cart.Props
	}
	type want struct {
		res *cart.Cart
		err error
	}
	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			args.ctx = prepare(args.ctx)

			got, err := s.cartService.Save(args.ctx, args.props)
			s.Equal(got, want.res)
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
			name: "product not found",
			prepare: func(ctx context.Context) context.Context {
				s.productRepo.EXPECT().
					Get(ctx, &product.Filter{
						ID: product.IDFilter{
							Eq: product.IDs{1},
						},
					}).
					Return(nil, errorx.ErrNotFound)
				return ctx
			},
			args: args{
				ctx: context.Background(),
				props: &cart.Props{
					ProductID: 1,
					Amount:    1,
				},
			},
			want: want{
				res: nil,
				err: errorx.ErrNotFound,
			},
		},
		{
			name: "success",
			prepare: func(ctx context.Context) context.Context {
				s.productRepo.EXPECT().
					Get(ctx, &product.Filter{
						ID: product.IDFilter{
							Eq: product.IDs{1},
						},
					}).
					Return(&product.Product{
						ID: 1,
						Props: &product.Props{
							Amount: 10,
						},
					}, nil)
				s.cartRepo.EXPECT().
					Save(ctx, &cart.Props{
						ProductID: 1,
						Amount:    1,
					}).
					Return(&cart.Cart{
						ID:     1,
						UserID: 1,
						Props: &cart.Props{
							ProductID: 1,
							Amount:    1,
						},
					}, nil)
				return ctx
			},
			args: args{
				ctx: context.Background(),
				props: &cart.Props{
					ProductID: 1,
					Amount:    1,
				},
			},
			want: want{
				res: &cart.Cart{
					ID:     1,
					UserID: 1,
					Props: &cart.Props{
						ProductID: 1,
						Amount:    1,
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}
