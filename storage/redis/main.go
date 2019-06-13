package redis

import (
	"github.com/go-redis/redis"
	"strings"
	"test_task/config"
	"test_task/logger"
	"test_task/utils"
)

var log = logger.GetLogger("redis", "storage")

var addr = []string{config.Configuration.Storage.Redis.Host, config.Configuration.Storage.Redis.Port}

var Client = redis.NewClient(&redis.Options{
	DB: 0,  // use default DB
	Addr: strings.Join(addr, ":"),
	Password: config.Configuration.Storage.Redis.Password, // no password set
})

func CheckState() {
	_, err := Client.Ping().Result()
	utils.HandleError(err, log)
	log.Info("Redis connection OK")
}