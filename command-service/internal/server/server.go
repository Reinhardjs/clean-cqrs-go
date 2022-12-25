package server

import (
	"github.com/Reinhardjs/golang-alpha-indo-soft/command-service/config"
	pkgElasticSearch "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/elasticsearch"
	pkgNats "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/nats"
	pkgPostgres "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/postgres"
	pkgRedis "github.com/Reinhardjs/golang-alpha-indo-soft/pkg/redis"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/gomodule/redigo/redis"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type server struct {
	cfg                 *config.Config
	PostgresConn        *gorm.DB
	RedisConn           redis.Conn
	ElasticSearchClient *elastic.Client
	NatsConn            *nats.Conn
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run() error {

	postgresConn, err := pkgPostgres.NewPostgresConn(s.cfg.Postgres)
	if err != nil {
		return errors.Wrap(err, "postgresConn")
	}
	s.PostgresConn = postgresConn

	redisConn, err := pkgRedis.NewRedisConn(s.cfg.Redis)
	if err != nil {
		return errors.Wrap(err, "redisConn")
	}
	s.RedisConn = redisConn

	natsConn, err := pkgNats.NewNatsEventStoreConn(s.cfg.Nats)
	if err != nil {
		return errors.Wrap(err, "natsConn")
	}
	s.NatsConn = natsConn

	elasticSearchClient, err := pkgElasticSearch.NewElasticSearchClient(s.cfg.ElasticSearch)
	if err != nil {
		return errors.Wrap(err, "elasticSearchClient")
	}
	s.ElasticSearchClient = elasticSearchClient

	return nil
}
