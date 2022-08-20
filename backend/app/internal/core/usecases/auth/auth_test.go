package auth_test

import (
	"github.com/blackPavlin/shop/app/internal/config"
	"github.com/blackPavlin/shop/app/internal/core/usecases/auth"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// AuthUseCaseBaseSuite
type AuthUseCaseBaseSuite struct {
	suite.Suite
	Ctrl *gomock.Controller

	userService        *auth.MockUserService
	cartService        *auth.MockCartService
	transactionService *auth.MockTransactionService
	authConfig         config.AuthConfig
}

// SetupTest
func (s *AuthUseCaseBaseSuite) SetupTest() {
	s.Ctrl = gomock.NewController(s.T())

	s.userService = auth.NewMockUserService(s.Ctrl)
	s.cartService = auth.NewMockCartService(s.Ctrl)
	s.transactionService = auth.NewMockTransactionService(s.Ctrl)

	s.authConfig = config.AuthConfig{}
}

// AuthUseCaseSuite
type AuthUseCaseSuite struct {
	AuthUseCaseBaseSuite

	authUseCase auth.AuthUseCase
}

// SetupTest
func (s *AuthUseCaseSuite) SetupTest() {
	s.AuthUseCaseBaseSuite.SetupTest()

	s.authUseCase = auth.NewAuthUseCase(
		s.userService,
		s.cartService,
		s.transactionService,
		s.authConfig,
	)
}

// TestHandler_Login
func (s *AuthUseCaseSuite) TestHandler_Login() {}

// TestHandler_Signup
func (s *AuthUseCaseSuite) TestHandler_Signup() {}
