package dto

type GetProductQueryDTO struct {
	Category string `form:"catergory"`
	Limit    int64  `form:"limit"`
	Offset   int64  `form:"offset"`
}

type CreateProductDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Quantity    int64  `json:"quantity" binding:"required"`
}
