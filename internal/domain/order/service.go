package order

// OrderService
type OrderService interface{}

// OrderUseCase
type OrderUseCase struct {
	orderRepo Repository
}

// NewOrderUseCase
func NewOrderUseCase(orderRepo Repository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}
