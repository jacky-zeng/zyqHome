package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"zyqHome/backProject/internal/config"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis(cfg *config.Config) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisAuth,
		DB:       0,
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect Redis: %v", err)
	}

	log.Println("Redis connected successfully")
}
