package main

import (
	"fmt"
	"os"
	"product-service/internal/config"
	"product-service/internal/middleware"
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

	router.Use(func(c *gin.Context) {
		fmt.Printf("[DEBUG] Request Masuk: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
		fmt.Printf("[DEBUG] Response Selesai: %d\n", c.Writer.Status())
	})

	router.GET("/products", productHdr.GetAllProducts)

	seller := router.Group("/seller")
	seller.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET"), "SELLER"))
	{
		seller.POST("/products", productHdr.AddProducts)
	}

	router.Run("localhost:8081")
}
