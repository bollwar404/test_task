package places

import (
	"test_task/models/place"

	"github.com/gin-gonic/gin"
)

func FetchPlace(c *gin.Context) {
	identifier := c.Param("identifier")
	locale := c.DefaultQuery("locale", "en")
	result := place.FetchPlace(identifier, locale)
	c.JSON(200, result)
}
