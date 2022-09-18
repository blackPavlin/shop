package order

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "order"

// OrderRepository
type OrderRepository interface{}
