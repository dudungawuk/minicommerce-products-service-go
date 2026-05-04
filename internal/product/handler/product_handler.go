package handler

import (
	"fmt"
	"net/http"
	"product-service/internal/product/dto"
	"product-service/internal/product/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAllProducts(c *gin.Context)
	AddProducts(c *gin.Context)
}

type productHandler struct {
	s service.ProductService
}

func NewProductHandler(s service.ProductService) ProductHandler {
	return &productHandler{
		s: s,
	}
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Print(ctx)
	var products, err = h.s.GetAllProducts(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, products)
}

func (h *productHandler) AddProducts(c *gin.Context) {
	var req dto.CreateProductRequest
	ctx := c.Request.Context()

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, dto.WebResponse{
			Code:    400,
			Status:  "BAD REQUEST",
			Message: "Input tidak valid",
			Data:    err.Error(),
		})
		return
	}

	err := h.s.AddProducts(ctx, req)
	if err != nil {
		resp := dto.WebResponse{
			Code:    500,
			Status:  "ERR",
			Message: "Unexpected Error",
			Data:    err.Error(),
		}
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}
	resp := dto.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "data added to database",
		Data:    err,
	}
	c.IndentedJSON(http.StatusOK, resp)
}
