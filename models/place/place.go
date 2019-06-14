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
	result := make(Places, 0)

	fromCache := cacheStore.Get(identifier, locale)
	if fromCache == "" {
		log.Info("No cache by identifier ", identifier, "_", locale)
		data := aviasales.GetPlaces(identifier, locale)
		for _, item := range data {
			if item.Type == "city" {
				result = append(result, Place{item.Code, item.CountryName, item.Name})
			} else {
				result = append(result, Place{item.Code, item.CityName, item.Name})
			}
		}

		if len(result) != 0 {
			cacheStore.Set(identifier, locale, string(utils.ToJSON(result, log)))
		}
	} else {
		utils.FromJSON(fromCache, &result, log)
	}

	return result
}
