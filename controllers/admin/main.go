package admin

import (
	"github.com/gin-gonic/gin"
	"test_task/models/admin"
)

func DropCache(c *gin.Context) {
	admin.DropCache()
	c.JSON(200, gin.H{
		"result": "ok",
	})
}

