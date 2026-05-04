package main

import (
	"product-service/internal/config"
	"product-service/internal/product/handler"
	"product-service/internal/product/repository"
	"product-service/internal/product/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productHdr := handler.NewProductHandler(productSvc)

	router := gin.Default()
	router.GET("/products", productHdr.GetAllProducts)
	router.POST("/products", productHdr.AddProducts)
	router.Run("localhost:8080")
}
