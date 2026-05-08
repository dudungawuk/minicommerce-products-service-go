package dto

type CreateProductRequest struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Stock    int     `json:"stock" binding:"required"`
	Category string  `json:"category" binding:"required"`
}

type CreateProductResponse struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Stock    int     `json:"stock" binding:"required"`
	Category string  `json:"category" binding:"required"`
}
