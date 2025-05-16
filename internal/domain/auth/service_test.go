package auth_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	"github.com/blackPavlin/shop/internal/config"
	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/testutilx"
)

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceSuite))
}

type AuthServiceBaseSuite struct {
	testutilx.BaseSuite
	userRepo *user.MockRepository
	config   *config.AuthConfig
	logger   *zap.Logger
}

func (s *AuthServiceBaseSuite) SetupTest() {
	s.BaseSuite.SetupTest()
	s.userRepo = user.NewMockRepository(s.Ctrl)
	s.config = &config.AuthConfig{
		SigningKey: "random",
		ExpiresIn:  60 * 60 * 1000,
	}
	s.logger = zaptest.NewLogger(s.T())
}

type AuthServiceSuite struct {
	AuthServiceBaseSuite
	authService auth.Service
}

func (s *AuthServiceSuite) SetupTest() {
	s.AuthServiceBaseSuite.SetupTest()
	s.authService = auth.NewUseCase(s.logger, s.config, s.userRepo)
}

func (s *AuthServiceSuite) TestSignup() {
	type args struct {
		ctx   context.Context
		props *auth.SignupProps
	}
	type want struct {
		res string
		err error
	}
	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			args.ctx = prepare(args.ctx)

			got, err := s.authService.Signup(args.ctx, args.props)
			s.Equal(want.res, got)
			s.ErrorIs(err, want.err)
		}
	}
	signupProps1 := &auth.SignupProps{
		Email:    gofakeit.Email(),
		Name:     gofakeit.Name(),
		Phone:    gofakeit.Phone(),
		Password: gofakeit.Password(true, true, true, true, true, 4),
	}
	tests := []struct {
		name    string
		prepare func(ctx context.Context) context.Context
		args    args
		want    want
	}{
		{
			name: "user already exists",
			prepare: func(ctx context.Context) context.Context {
				s.userRepo.EXPECT().
					Exist(ctx, &user.Filter{
						Email: user.EmailFilter{Eq: []string{signupProps1.Email}},
					}).
					Return(true, nil)
				return ctx
			},
			args: args{
				ctx:   context.Background(),
				props: signupProps1,
			},
			want: want{
				res: "",
				err: errorx.ErrAlreadyExists,
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}

func (s *AuthServiceSuite) TestLogin() {
	type args struct {
		ctx   context.Context
		props *auth.LoginProps
	}
	type want struct {
		res string
		err error
	}
	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			args.ctx = prepare(args.ctx)

			got, err := s.authService.Login(args.ctx, args.props)
			s.Equal(want.res, got)
			s.ErrorIs(err, want.err)
		}
	}
	testUser := &user.User{
		ID:   1,
		Role: user.RoleUser,
		Props: &user.Props{
			Email:    "user@user.com",
			Name:     "testUser",
			Phone:    "80000000000",
			Password: "$2a$14$VftkClux8ptrh/lZBU99ROtOs5vq1rLFEpbOxuqrbqPUOSHGCEVA.",
		},
	}
	invalidLoginProps := &auth.LoginProps{
		Email:    "some@email.com",
		Password: "password123",
	}
	invalidPasswordProps := &auth.LoginProps{
		Email:    testUser.Props.Email,
		Password: "password123",
	}
	tests := []struct {
		name    string
		prepare func(ctx context.Context) context.Context
		args    args
		want    want
	}{
		{
			name: "invalid login",
			prepare: func(ctx context.Context) context.Context {
				s.userRepo.EXPECT().
					Get(ctx, &user.Filter{
						Email: user.EmailFilter{Eq: []string{invalidLoginProps.Email}},
					}).
					Return(nil, errorx.ErrNotFound)
				return ctx
			},
			args: args{
				ctx:   context.Background(),
				props: invalidLoginProps,
			},
			want: want{
				res: "",
				err: errorx.ErrInvalidLoginOrPassword,
			},
		},
		{
			name: "invalid password",
			prepare: func(ctx context.Context) context.Context {
				s.userRepo.EXPECT().
					Get(ctx, &user.Filter{
						Email: user.EmailFilter{Eq: []string{invalidPasswordProps.Email}},
					}).
					Return(testUser, nil)
				return ctx
			},
			args: args{
				ctx:   context.Background(),
				props: invalidPasswordProps,
			},
			want: want{
				res: "",
				err: errorx.ErrInvalidLoginOrPassword,
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}
