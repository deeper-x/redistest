package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// DB redis client struct
type DB struct {
	Client *redis.Client
	Ctx    context.Context
}

func main() {
	db := NewDB()

	err := db.SetValue("demo", "value", time.Duration(0)*time.Second)
	if err != nil {
		log.Panic(err)
	}

}

// NewDB return myclient instance
func NewDB() DB {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	return DB{
		Client: cli,
		Ctx:    ctx,
	}
}

// SetValue is a redis set wrapper
func (db DB) SetValue(key string, value interface{}, expiration time.Duration) error {
	err := db.Client.Set(db.Ctx, key, value, expiration).Err()

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
