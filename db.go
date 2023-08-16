package main

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase(address, password string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}
