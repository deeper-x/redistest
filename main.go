package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// DBManager is interface for redis db
type DBManager interface {
	SetValue(key string, value interface{}, expiration time.Duration) error
}

// DB redis client struct
type DB struct {
	Client *redis.Client
	Ctx    context.Context
}

func main() {
	db := NewDB()

	DoSomething(db, "demo", "value")
	log.Println(db)
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

// DoSomething is a mocked function
func DoSomething(db DBManager, key, value string) {
	err := db.SetValue(key, value, time.Duration(0)*time.Second)

	if err != nil {
		panic(err)
	}
}
