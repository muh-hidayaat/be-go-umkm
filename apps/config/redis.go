// config/redis.go

package config

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Ctx = context.Background()

func CreateRedisClient() *redis.Client {
	redisAddr := viper.GetString("REDIS_ADDR")
	redisPassword := viper.GetString("REDIS_PASSWORD")
	redisDB := viper.GetString("REDIS_DB")

	options := &redis.Options{
		Addr:     redisAddr,
		Password: redisPassword, // Kosongkan jika tidak ada password
	}

	if redisDB != "" {
		db, err := strconv.Atoi(redisDB)
		if err == nil {
			options.DB = db
		} else {
			log.Println("Invalid REDIS_DB value, using default DB 0")
		}
	}

	rdb := redis.NewClient(options)

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Redis connected ... !")

	return rdb
}
