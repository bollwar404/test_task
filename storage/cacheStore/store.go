package cacheStore

import (
	"test_task/config"
	"test_task/logger"
	"test_task/storage"
	"test_task/utils"

	"time"
)

var log = logger.GetLogger("cacheStore", "storage")

var redisClient = storage.RedisClient

const cachePrefix = "cache_places_"

func getKey(id, locale string) string {
	return cachePrefix + id + "_" + locale
}

func Set(id, locale string, value string)  {
	key := getKey(id, locale)
	log.Info("set cache by key ", key)
	err := redisClient.Set(key, value, config.Configuration.Cache.Places.Ttl * time.Second * 60 * 60 * 24).Err()
	if err != nil {
		log.Error(err.Error())
	}
}

func Get(id, locale string) string{
	key := getKey(id, locale)
	log.Info("get cache by key ", key)
	val, err := redisClient.Get(key).Result()
	if err != nil && err.Error() != "redis: nil" {
		log.Error(err.Error())
	}
	return val
}

func Drop() {
	log.Info("drop cache ", cachePrefix)
	var cursor uint64
	for {
		keys, cursor, err := redisClient.Scan(cursor, cachePrefix + "*", 10).Result()
		utils.HandleError(err, log)
		for _, key := range keys {
			err := redisClient.Del(key).Err()
			if err != nil {
				log.Error(err.Error())
			}
		}
		if cursor == 0 {
			break
		}
	}
}
