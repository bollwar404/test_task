package place

import (
	"test_task/logger"
	"test_task/requests/aviasales"
	"test_task/storage/cacheStore"
	"test_task/utils"
)

var log = logger.GetLogger("place", "models")

type Place struct {
	Slug string `json:"slug"`
	Subtitle string `json:"subtitle"`
	Title string `json:"title"`
}

type Places []Place

func FetchPlace(identifier, locale string) Places {
	var result Places

	cacheKey := identifier + "_" + locale

	fromCache := cacheStore.Get(cacheKey)
	if fromCache != "" {
		log.Info("From cache by identifier ", cacheKey)
		utils.FromJSON(fromCache, &result, log)
		return result
	}

	data := aviasales.GetPlaces(identifier, locale)
	for _, item := range data {
		if item.Type == "city" {
			result = append(result, Place{item.Code, item.CountryName, item.Name})
		} else {
			result = append(result, Place{item.Code, item.CityName, item.Name})
		}
	}

	cacheStore.Set(cacheKey, string(utils.ToJSON(result, log)))

	return result
}
