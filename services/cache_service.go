package services

import (
	"context"
	"time"

	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var cache *redis.Client

func StartCache() {
	cacheConf := config.GetConfig().CACHE

	addr := cacheConf.Host + ":" + cacheConf.Port
	url := "redis://" + cacheConf.User + ":" + cacheConf.Password + "@" + addr

	newCache, err := redis.ParseURL(url)
	if err != nil {
		logger.Error("error parsing cache URL", err)
	}

	cache = redis.NewClient(newCache)

	if err := cache.Ping(ctx).Err(); err != nil {
		logger.Error("error connecting to cache", err)
		return
	}

	logger.Info("connected to cache at port: " + cacheConf.Port)
}

func GetCache(key string) string {
	val, err := cache.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

func SetCache(key string, value string) string {
	var ttl = config.GetConfig().CACHE.ExpiresIn

	err := cache.Set(ctx, key, value, ttl*time.Second).Err()
	if err != nil {
		logger.Error("error setting cache", err)
	}

	return value
}

func DeleteCache(key string) {
	err := cache.Del(ctx, key).Err()
	if err != nil {
		logger.Error("error deleting cache", err)
	}
}
