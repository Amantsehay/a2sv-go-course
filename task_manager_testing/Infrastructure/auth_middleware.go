package Infrastructure

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := ParseToken(authParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("user")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT"})
			c.Abort()
			return
		}

		userClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT claims"})
			c.Abort()
			return
		}

		if userClaims["role"] != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
