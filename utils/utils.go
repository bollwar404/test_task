package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(err error, log *log.Entry) {
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

func FromJSON(jsonString string, result interface{}, log *log.Entry) {
	err := json.Unmarshal([]byte(jsonString), &result)
	HandleError(err, log)
}

func ToJSON(data interface{}, log *log.Entry) []byte {
	var jsonData []byte
	jsonData, err := json.Marshal(data)
	HandleError(err, log)

	return jsonData
}

func Validate(field, tag string, c *gin.Context)  {

}
