package order

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "order"

// Repository
type Repository interface{}
