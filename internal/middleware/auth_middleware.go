package middleware

import (
	"net/http"
	"product-service/internal/product/dto"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secretKey string, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.WebResponse{
				Code:    401,
				Status:  "UNAUTHORIZED",
				Message: "Token ga ada atau format salah (pake Bearer)",
			})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.WebResponse{
				Code:    401,
				Status:  "UNAUTHORIZED",
				Message: "Token palsu atau udah expired",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.WebResponse{
				Code:    401,
				Status:  "UNAUTHORIZED",
				Message: "Gagal ambil claims",
			})
			return
		}

		role, ok := claims["role"].(string)
		if !ok || (requiredRole != "" && role != requiredRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.WebResponse{
				Code:    403,
				Status:  "FORBIDDEN",
				Message: "Role lo ga cukup buat akses ini",
			})
			return
		}

		c.Set("username", claims["sub"])
		c.Set("role", role)

		c.Next()
	}
}
