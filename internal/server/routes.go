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
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	apiGroup := r.Group("/api/v1")

	userRoutes := apiGroup.Group("/users")
	userRoutes.Use(middleware.BasicAuthMiddleware())
	{
		userRoutes.POST("/create", s.userHandler.FindOrCreateUser)
		userRoutes.GET("/:id", s.userHandler.GetUserById)
	}

	boardRoutes := apiGroup.Group("/board")
	boardRoutes.Use(middleware.BasicAuthMiddleware())
	{
		postRoutes := boardRoutes.Group("/posts")
		{
			postRoutes.GET("", s.postHandler.ListPosts)
			postRoutes.POST("", s.postHandler.CreatePost)
			postRoutes.GET("/:id", s.postHandler.GetPost)
			postRoutes.DELETE("/:id", s.postHandler.DeletePost)
			postRoutes.POST("/:id/upvote", s.postHandler.UpvotePost)
			postRoutes.POST("/:id/downvote", s.postHandler.DownvotePost)
			postRoutes.POST("/:id/report", s.postHandler.ReportPost)

			postRoutes.GET("/:id/comments", s.commentHandler.ListComments)
			postRoutes.POST("/:id/comments", s.commentHandler.AddComment)
		}

		commentRoutes := boardRoutes.Group("/comments")
		{
			commentRoutes.DELETE("/:id", s.commentHandler.DeleteComment)
			commentRoutes.PUT("/:id", s.commentHandler.UpdateComment)
			commentRoutes.POST("/:id/upvote", s.commentHandler.UpvoteComment)
			commentRoutes.POST("/:id/downvote", s.commentHandler.DownvoteComment)
		}

		feedRoutes := boardRoutes.Group("/feed")
		{
			feedRoutes.GET("", s.feedHandler.GetFeed)
		}
	}

	return r
}
