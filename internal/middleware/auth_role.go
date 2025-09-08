package middleware

import (
	"chin_server/internal/model"

	"github.com/gin-gonic/gin"
)

// AdminOnly middleware: chỉ cho phép admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("id").(*model.User)
		if !ok || user.Role != "admin" {
			c.JSON(403, gin.H{"error": "forbidden: admin only"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// UserOnly middleware: chỉ cho phép user thường
func UserOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("id").(*model.User)
		if !ok || user.Role != "user" {
			c.JSON(403, gin.H{"error": "forbidden: user only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
