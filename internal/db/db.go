package db

import (
	"fmt"

	"github.com/yourusername/nebula-grid/internal/models"
	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Open creates a gorm DB connection based on config and runs auto-migrations.
func Open(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	if cfg.DBDriver == "postgres" {
		if cfg.PostgresDSN == "" {
			return nil, fmt.Errorf("POSTGRES_DSN is required for postgres driver")
		}
		dialector = postgres.Open(cfg.PostgresDSN)
	} else {
		// default to sqlite
		dialector = sqlite.Open("nebula.db")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.Node{}, &models.Service{}); err != nil {
		logger.Log.Errorf("auto-migrate failed: %v", err)
		return nil, err
	}

	return db, nil
}
