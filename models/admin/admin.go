package admin

import (
	"test_task/logger"
	"test_task/storage/cacheStore"
)

var log = logger.GetLogger("admin", "models")

func DropCache() {
	log.Info("Drop Cache start")
	cacheStore.Drop()
	log.Info("Drop Cache fin")
}
