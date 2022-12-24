package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type Config struct {
	Host string
	Port string
}

func NewRedisConn(cfg *Config) (redis.Conn, error) {
	redisUri := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	conn, err := redis.Dial("tcp", redisUri)

	if err != nil {
		return nil, errors.Wrap(err, "redis.ConnectConfig")
	}

	return conn, nil
}
