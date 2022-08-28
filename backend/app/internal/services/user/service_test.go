package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/blackPavlin/shop/app/internal/entities"
	"github.com/blackPavlin/shop/app/internal/services/user"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type UserServiceBaseSuite struct {
	suite.Suite
	Ctrl *gomock.Controller

	userRepository *user.MockUserRepository
}

func (s *UserServiceBaseSuite) SetupTest() {
	s.Ctrl = gomock.NewController(s.T())
	s.userRepository = user.NewMockUserRepository(s.Ctrl)
}

type UserServiceSuite struct {
	UserServiceBaseSuite
	userService user.UserService
}

func (s *UserServiceSuite) SetupTest() {
	s.UserServiceBaseSuite.SetupTest()
	s.userService = user.NewUserService(s.userRepository)
}

func (s *UserServiceSuite) TestHandler_CreateUser() {
	type args struct {
		ctx  context.Context
		user *entities.User
	}
	type want struct {
		user entities.UserID
		err  error
	}

	test := func(prepare func(ctx context.Context) context.Context, args args, want want) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			args.ctx = prepare(args.ctx)

			got, err := s.userService.Create(args.ctx, args.user)
			s.Equal(want.user, got)
			s.True(errors.Is(err, want.err))
		}
	}

	user := &entities.User{
		Name:  gofakeit.Name(),
		Phone: gofakeit.Phone(),
		Email: gofakeit.Email(),
		Role:  entities.RoleUser,
	}

	userID := entities.UserID(primitive.NewObjectID())

	tests := []struct {
		name    string
		prepare func(ctx context.Context) context.Context
		args    args
		want    want
	}{
		{
			name: "User successfully created",
			prepare: func(ctx context.Context) context.Context {
				s.userRepository.EXPECT().Create(ctx, user).Return(userID, nil)

				return ctx
			},
			args: args{
				ctx:  context.Background(),
				user: user,
			},
			want: want{
				user: userID,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(tt.name, test(tt.prepare, tt.args, tt.want))
	}
}
