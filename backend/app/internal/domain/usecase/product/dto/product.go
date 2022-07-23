package dto

type CreateProductDTO struct {
	Name        string
	Description string
	Category    string
	Quantity    int64
}

type GetProductDTO struct {
	Category string
	Limit    int64
	Offset   int64
}

type GetProductOutDTO struct {
	ID          string
	Name        string
	Description string
	Category    string
	Quantity    int64
}
