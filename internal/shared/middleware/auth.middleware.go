package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}