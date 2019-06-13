package cacheStore

import (
	"test_task/logger"
	"test_task/storage"
	"test_task/utils"
)

var log = logger.GetLogger("cacheStore", "storage")

var redisClient = storage.RedisClient

const cachePrefix = "cache_"

func Set(key string, value string)  {
	log.Info("get cache by key ", key)
	err := redisClient.Set(cachePrefix + key, value, 0).Err()
	utils.HandleError(err, log)
}

func Get(key string) string{
	log.Info("set cache by key ", key)
	val, err := redisClient.Get(cachePrefix + key).Result()
	if err != nil && err.Error() != "redis: nil" {
		log.Error(err.Error())
		panic(err)
	}
	return val
}
