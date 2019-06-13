package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"test_task/config"
	"test_task/controllers/places"
	"test_task/logger"
	"test_task/storage"
)

var log = logger.GetLogger("main", "root")

func main() {
	log.Info("Start server creation")
	storage.CheckDBConnections()
	router := gin.Default()
	v1 := router.Group("/api/v1")

	placesGroup := v1.Group("/places")
	{
		placesGroup.GET("/:identifier", places.FetchPlace)
	}

	serverAddress := strings.Join([]string{config.Configuration.Server.Host, config.Configuration.Server.Port}, ":")
	log.Info("Server listening ", serverAddress)
	router.Run(serverAddress)
}



