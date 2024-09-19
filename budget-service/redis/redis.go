package redis

import (
    "log"
    "os"

    "github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")

    client := redis.NewClient(&redis.Options{
        Addr: redisHost + ":" + redisPort,
    })

    log.Println("Connected to Redis")
    return client
}
