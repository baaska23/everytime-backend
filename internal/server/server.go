package server

import (
	"everytime-backend/internal/shared/database"
	"everytime-backend/internal/users"
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
	userHandler *users.Handler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbManager, err := database.NewDBManager()
	if err != nil {
		log.Fatalf("Failed to init database manager: %v", err)
	}

	userRepo := users.NewRepository(dbManager.Everytime)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

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
