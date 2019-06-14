package cacheStore

import (
	"test_task/logger"
	"test_task/storage"
)

var log = logger.GetLogger("cacheStore", "storage")

var redisClient = storage.RedisClient

const cachePrefix = "cache_places"

func Set(key string, value string)  {
	log.Info("set cache by key ", key)
	err := redisClient.HSet(cachePrefix, key, value).Err()
	if err != nil {
		log.Error(err.Error())
	}
}

func Get(key string) string{
	log.Info("get cache by key ", key)
	val, err := redisClient.HGet(cachePrefix, key).Result()
	if err != nil && err.Error() != "redis: nil" {
		log.Error(err.Error())
	}
	return val
}

func Drop() {
	log.Info("drop cache ", cachePrefix)
	err := redisClient.Del(cachePrefix).Err()
	if err != nil {
		log.Error(err.Error())
	}
}
