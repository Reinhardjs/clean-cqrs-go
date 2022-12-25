package config

import (
	"os"

	"github.com/Reinhardjs/golang-alpha-indo-soft/pkg/constants"
	"github.com/Reinhardjs/golang-alpha-indo-soft/pkg/elasticsearch"
	"github.com/Reinhardjs/golang-alpha-indo-soft/pkg/nats"
	"github.com/Reinhardjs/golang-alpha-indo-soft/pkg/postgres"
	"github.com/Reinhardjs/golang-alpha-indo-soft/pkg/redis"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Postgres      *postgres.Config
	Redis         *redis.Config
	ElasticSearch *elasticsearch.Config
	Nats          *nats.Config
}

func InitConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "godotenv.Load")
	}

	cfg := &Config{}

	postgresHost := os.Getenv(constants.PostgresHost)
	postgresDBName := os.Getenv(constants.PostgresDBName)
	postgresUser := os.Getenv(constants.PostgresUser)
	postgresPassword := os.Getenv(constants.PostgresPassword)
	postgresSSLMode := os.Getenv(constants.PostgresSSLMode)
	cfg.Postgres = &postgres.Config{
		Host:     postgresHost,
		DBName:   postgresDBName,
		User:     postgresUser,
		Password: postgresPassword,
		SSLMode:  postgresSSLMode,
	}

	redisHost := os.Getenv(constants.RedisHost)
	redisPort := os.Getenv(constants.RedisPort)
	cfg.Redis = &redis.Config{
		Host: redisHost,
		Port: redisPort,
	}

	natsAddress := os.Getenv(constants.NatsAddress)
	cfg.Nats = &nats.Config{
		Url: natsAddress,
	}

	elasticsearchAddress := os.Getenv(constants.ElasticsearchAddress)
	elasticsearchUsername := os.Getenv(constants.ElasticsearchUsername)
	elasticsearchPassword := os.Getenv(constants.ElasticsearchPassword)
	cfg.ElasticSearch = &elasticsearch.Config{
		Url:      elasticsearchAddress,
		Username: elasticsearchUsername,
		Password: elasticsearchPassword,
	}

	return cfg, nil
}
