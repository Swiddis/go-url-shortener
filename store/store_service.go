package store

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "time"
)

type StorageService struct {
    redisClient *redis.Client
}

var (
    storeService = &StorageService{}
    ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
    redisClient := redis.NewClient(&redis.Options{
        Addr: "redis:6379",
        Password: "YBVc5bXoiRlDLxYt9mXtvYPoxLRaDA7m96MBFDhd",
        DB: 0,
    })

    pong, err := redisClient.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("Error init Redis: %v", err))
    }

    fmt.Printf("\nRedis client started successfully: pong message = {%s}", pong)
    storeService.redisClient = redisClient
    return storeService
}

func SaveUrlMapping(shortUrl, originalUrl, userId string) {
    err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
    if err != nil {
        panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
    }
}

func RetrieveInitialUrl(shortUrl string) string {
    result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
    if err != nil {
        panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
    }
    return result
}
