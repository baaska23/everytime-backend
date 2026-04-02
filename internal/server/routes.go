package server

import (
	"everytime-backend/internal/shared/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	apiGroup := r.Group("/api/v1")

	userRoutes := apiGroup.Group("/users")
	userRoutes.Use(middleware.BasicAuthMiddleware())
	{
		userRoutes.POST("/create", s.userHandler.FindOrCreateUser)
		userRoutes.GET("/:id", s.userHandler.GetUserById)
	}

	return r
}
