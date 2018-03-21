package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

// DB is a wrapper interface, stub it out in tests of other packages
type DB interface {
	GetValue(key string) (interface{}, error)
	SetValue(item Item) error
	Close() error
	FlushAll() error
}

type RedisDB struct {
	*redis.Client
}

func Connect() DB {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if pong == "" || err != nil {
		log.Fatal(err)
	}

	return &RedisDB{Client: client}
}

func (db *RedisDB) GetValue(key string) (interface{}, error) {
	return nil, nil
}

func (db *RedisDB) SetValue(item Item) error {
	return nil
}

func (db *RedisDB) Close() error {
	return nil
}

func (db *RedisDB) FlushAll() error {
	return nil
}
