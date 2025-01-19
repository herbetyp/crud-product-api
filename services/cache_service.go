package services

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	config "github.com/herbetyp/crud-product-api/internal/configs"
	"github.com/herbetyp/crud-product-api/internal/configs/logger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var cache *redis.Client

func StartCache() {
	cacheConf := config.GetConfig().CACHE

	addr := cacheConf.Host + ":" + cacheConf.Port + "/" + strconv.Itoa(cacheConf.Db)
	url := "redis://" + ":" + cacheConf.Password + "@" + addr + "?protocol=3"

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

func SetCache(key string, i interface{}) interface{} {
	var ttl = config.GetConfig().CACHE.ExpiresIn

	cacheValue, err := json.Marshal(i)

	if err != nil {
		logger.Error("error marshalling cache", err)
	} else {
		err := cache.Set(ctx, key, string(cacheValue), ttl*time.Second).Err()
		if err != nil {
			logger.Error("error setting cache", err)
		}
	}

	return cacheValue
}

func GetCache(key string, i interface{}) string {
	cacheData, err := cache.Get(ctx, key).Result()

	if err != redis.Nil {
		logger.Error("error getting cache", err)
	}

	if cacheData != "" {
		err = json.Unmarshal([]byte(cacheData), i)
		if err != nil {
			logger.Error("error unmarshalling cache", err)
		}
	}

	return cacheData
}

func DeleteCache(key string, prefix string, flushall bool) {
	err := cache.Del(ctx, prefix+key).Err()
	if err != nil {
		logger.Error("error deleting cache", err)
	}

	if flushall {
		err := cache.Del(ctx, prefix+"all").Err()
		if err != nil {
			logger.Error("error flushing cache", err)
		}
	}
}
