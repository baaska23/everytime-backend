package middleware

import (
	"everytime-backend/internal/shared/apierror"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}

func BearerAuthMiddleware() gin.HandlerFunc {
	secret := os.Getenv("JWT_SECRET")

	return func(c *gin.Context) {
		if secret == "" {
			apierror.WriteGin(c, apierror.Internal(fmt.Errorf("JWT_SECRET is not configured")))
			return
		}

		authHeader := c.GetHeader("Authorization")
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			apierror.WriteGin(c, apierror.Unauthorized("missing or invalid authorization header"))
			return
		}

		rawToken := strings.TrimSpace(parts[1])
		parsed, err := jwt.Parse(rawToken, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil || !parsed.Valid {
			apierror.WriteGin(c, apierror.Unauthorized("invalid token"))
			return
		}

		claims, ok := parsed.Claims.(jwt.MapClaims)
		if !ok {
			apierror.WriteGin(c, apierror.Unauthorized("invalid token claims"))
			return
		}

		tokenType, _ := claims["type"].(string)
		if tokenType != "access" {
			apierror.WriteGin(c, apierror.Unauthorized("access token required"))
			return
		}

		subject, _ := claims["sub"].(string)
		email, _ := claims["email"].(string)
		if subject == "" || email == "" {
			apierror.WriteGin(c, apierror.Unauthorized("invalid token payload"))
			return
		}

		c.Set("user_id", subject)
		c.Set("email", email)
		c.Next()
	}
}
