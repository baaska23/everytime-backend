package server

import (
	"everytime-backend/internal/ads"
	"everytime-backend/internal/auth"
	"everytime-backend/internal/board"
	"everytime-backend/internal/shared/database"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port           int
	dbManager      *database.DBManager
	userHandler    *auth.Handler
	postHandler    *board.PostHandler
	commentHandler *board.CommentHandler
	feedHandler    *board.FeedHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbManager, err := database.NewDBManager()
	if err != nil {
		log.Fatalf("Failed to init database manager: %v", err)
	}

	if err := dbManager.Everytime.AutoMigrate(&auth.User{}, &auth.OTPRecord{}, &auth.RefreshToken{}); err != nil {
		log.Fatalf("Failed to migrate auth tables: %v", err)
	}

	userRepo := auth.NewRepository(dbManager.Everytime)
	userService := auth.NewService(userRepo)
	userHandler := auth.NewHandler(userService)

	postRepo := board.NewPostRepository(dbManager.Everytime)
	postService := board.NewPostService(postRepo)
	postHandler := board.NewPostHandler(postService)

	commentRepo := board.NewCommentRepository(dbManager.Everytime)
	commentService := board.NewCommentService(commentRepo, postRepo)
	commentHandler := board.NewCommentHandler(commentService)

	adRepo := ads.NewAdRepository(dbManager.Everytime)
	feedService := board.NewFeedService(postRepo, adRepo)
	feedHandler := board.NewFeedHandler(feedService)

	srv := &Server{
		port:           port,
		dbManager:      dbManager,
		userHandler:    userHandler,
		postHandler:    postHandler,
		commentHandler: commentHandler,
		feedHandler:    feedHandler,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.port),
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
