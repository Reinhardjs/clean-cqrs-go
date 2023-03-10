package postgres

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

const (
	maxConn           = 50
	healthCheckPeriod = 1 * time.Minute
	maxConnIdleTime   = 1 * time.Minute
	maxConnLifetime   = 3 * time.Minute
	minConns          = 10
	lazyConnect       = false
)

func NewPostgresConn(cfg *Config) (*gorm.DB, error) {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", cfg.Host, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return conn, nil
}
