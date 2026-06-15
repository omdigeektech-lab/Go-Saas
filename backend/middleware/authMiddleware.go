package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Basic placeholder for step 1
		c.Set("user_id", "dummy_user_id")
		c.Next()
	}
}
