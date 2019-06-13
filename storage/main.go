package storage

import (
	"test_task/logger"
	"test_task/storage/redis"
)

var RedisClient = redis.Client
var log = logger.GetLogger("main", "storage")

func CheckDBConnections()  {
	log.Info("Start check db connections")
	redis.CheckState()
}
