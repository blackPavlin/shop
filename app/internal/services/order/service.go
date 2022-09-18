package order

// OrderService
type OrderService struct {
	orderRepository OrderRepository
}

// NewOrderService
func NewOrderService(orderRepository OrderRepository) OrderService {
	return OrderService{
		orderRepository: orderRepository,
	}
}
