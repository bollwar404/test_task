package aviasales

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"test_task/config"
	"test_task/logger"
	"test_task/utils"
	"time"
)

var log = logger.GetLogger("aviasales", "requests")

type PlacesResponse []struct {
	Code            string      `json:"code"`
	Name            string      `json:"name"`
	Type         string   `json:"type"`
	CountryName  string   `json:"country_name"`
	CityName string `json:"city_name,omitempty"`
}


func GetPlaces(identifier, locale string) PlacesResponse {
	var data PlacesResponse

	timeout := time.Duration(config.Configuration.Requests.Aviasales.Places.Timeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	var buffer bytes.Buffer
	buffer.WriteString(config.Configuration.Requests.Aviasales.Places.Url)
	buffer.WriteString(url.QueryEscape(identifier))
	buffer.WriteString("&locale=")
	buffer.WriteString(url.QueryEscape(locale))

	uri := buffer.String()

	log.Info("Start send request to ", uri)
	resp, err := client.Get(uri)
	if err, ok := err.(net.Error); ok && err.Timeout() {
		log.Error("Request timed out")
		return data
	}
	utils.HandleError(err, log)
	defer resp.Body.Close()
	log.Info("request finished")

	body, err := ioutil.ReadAll(resp.Body)
	utils.HandleError(err, log)

	utils.FromJSON(string(body), &data, log)

	return data
}
