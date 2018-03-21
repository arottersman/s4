package main

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

const KeyNotFound = DBError("S4: Key Not Found")

type DBError string

func (e DBError) Error() string { return string(e) }

type DB interface {
	GetValue(key string) (string, error)
	SetValue(item Item) error
	Close() error
	FlushDB() error
}

type redisDB struct {
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

	return &redisDB{Client: client}
}

func (db *redisDB) GetValue(key string) (string, error) {
	value, err := db.Client.Get(key).Result()
	if err == redis.Nil {
		// Key does not exist
		return "", KeyNotFound
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func (db *redisDB) SetValue(item Item) error {
	noExpiration := time.Duration(0)
	err := db.Client.Set(item.Key, item.Value, noExpiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *redisDB) Close() error {
	return db.Client.Close()
}

func (db *redisDB) FlushDB() error {
	return db.Client.FlushDB().Err()
}
