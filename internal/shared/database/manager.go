package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	Everytime *gorm.DB
}

func NewDBManager() (*DBManager, error) {
	manager := &DBManager{}

	everytimeDB, err := initEverytimeDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Everytime database: %w", err)
	}
	manager.Everytime = everytimeDB

	return manager, nil
}

func initEverytimeDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("getting underlying sql.DB: %w", err)
	}

	var version string
	if err := sqlDB.QueryRow("SELECT version()").Scan(&version); err != nil {
		return nil, fmt.Errorf("querying database version: %w", err)
	}

	log.Println("Connected to:", version)

	return db, nil
}
