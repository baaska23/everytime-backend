package server

import (
	"everytime-backend/internal/shared/database"
	"everytime-backend/internal/auth"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Server struct {
	port        int
	dbManager   *database.DBManager
	userHandler *auth.Handler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbManager, err := database.NewDBManager()
	if err != nil {
		log.Fatalf("Failed to init database manager: %v", err)
	}

	userRepo := auth.NewRepository(dbManager.Everytime)
	userService := auth.NewService(userRepo)
	userHandler := auth.NewHandler(userService)

	srv := &Server{
		port:        port,
		dbManager:   dbManager,
		userHandler: userHandler,
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
