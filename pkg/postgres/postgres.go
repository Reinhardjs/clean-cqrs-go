package postgres

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

type Config struct {
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbName"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  bool   `yaml:"sslMode"`
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
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", cfg.Host, cfg.User, cfg.DBName, cfg.sslMode, cfg.Password)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		return nil, errors.Wrap(err, "pgx.ConnectConfig")
	}

	return conn, nil
}
