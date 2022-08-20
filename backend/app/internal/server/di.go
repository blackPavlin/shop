package server

import (
	"github.com/blackPavlin/shop/app/internal/core/services"
	"github.com/blackPavlin/shop/app/internal/core/usecases"
	"github.com/blackPavlin/shop/app/internal/core/usecases/auth"
	"github.com/blackPavlin/shop/app/internal/core/usecases/user"
	"github.com/blackPavlin/shop/app/internal/repositories"
)

// initRepositories
func (s *Server) initRepositories() repositories.Repositories {
	return repositories.NewRepositories(s.mongo)
}

// initServices
func (s *Server) initServices() services.Services {
	return services.NewServices(
		s.repositories.UserRepository,
		s.repositories.CartRepository,
		s.repositories.AddressRepository,
		s.repositories.ProductRepository,
		s.repositories.TransactionRepository,
	)
}

// initUseCases
func (s *Server) initUseCases() usecases.UseCases {
	return usecases.UseCases{
		// AuthUseCase
		AuthUseCase: auth.NewAuthUseCase(
			s.services.UserService,
			s.services.CartService,
			s.services.TransactionService,
			s.config.Auth,
		),
		// UserUseCase
		UserUseCase: user.NewUserUseCase(
			s.services.UserService,
		),
	}
}
