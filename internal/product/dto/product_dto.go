package dto

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
