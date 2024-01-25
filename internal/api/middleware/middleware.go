package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token := bearerToken[7:]

		claims, err := jwt.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
