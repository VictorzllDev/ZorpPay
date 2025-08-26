package middleware

import (
	"net/http"
	"strings"

	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService security.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Necessary authorization token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format. Use Bearer Token"})
			c.Abort()
			return
		}

		userID, err := jwtService.GetUserIDFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided or invalid"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
