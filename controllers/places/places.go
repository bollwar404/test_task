package places

import (
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"test_task/logger"
	"test_task/models/place"

	"github.com/gin-gonic/gin"
)

var log = logger.GetLogger("1", "2")


func FetchPlace(c *gin.Context) {
	identifier := c.Param("identifier")
	locale := c.DefaultQuery("locale", "en")

	validate := validator.New()
	errs := validate.Var(identifier, "gte=1,lte=100,required")
	if errs != nil {
		log.Error("identifier ", identifier, errs)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation error",
		})
		return
	}
	errs = validate.Var(locale, "gt=0,lte=4,required,oneof=en ru")
	if errs != nil {
		log.Error("locale ", locale, errs)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation error",
		})
		return
	}

	result := place.FetchPlace(identifier, locale)
	if len(result) == 0 {
		c.JSON(200, []string{})
	} else {
		c.JSON(200, result)
	}
}
