package image

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "image"

// Service
type Service interface{}

// UseCase
type UseCase struct{}

// NewUseCase
func NewUseCase() *UseCase {
	return &UseCase{}
}
