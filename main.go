package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"strings"
	"test_task/config"
	"test_task/controllers/admin"
	"test_task/controllers/places"
	"test_task/logger"
	"test_task/storage"
	"test_task/utils"
)

var log = logger.GetLogger("main", "root")

func initAdminRoutes(group *gin.RouterGroup) {
	group.GET("/drop", admin.DropCache)
}

func initPlacesRoutes(group *gin.RouterGroup) {
	group.GET("/:identifier", places.FetchPlace)
}

func main() {
	log.Info("Start server creation")

	storage.CheckDBConnections()

	router := gin.New()
	router.Use(ginlogrus.Logger(logrus.New()), gin.Recovery())

	v1 := router.Group("/api/v1")

	placesGroup := v1.Group("/places")
	initPlacesRoutes(placesGroup)

	adminGroup := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	initAdminRoutes(adminGroup)

	serverAddress := strings.Join([]string{config.Configuration.Server.Host, config.Configuration.Server.Port}, ":")
	log.Info("Server listening ", serverAddress)
	err := router.Run(serverAddress)

	utils.HandleError(err, log)
}
