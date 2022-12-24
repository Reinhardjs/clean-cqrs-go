package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func NewPostgresConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		return nil, errors.Wrap(err, "redis.ConnectConfig")
	}

	return conn, nil
}
