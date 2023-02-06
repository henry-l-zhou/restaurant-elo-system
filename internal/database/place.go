package database

import (
	"fmt"
	"net/url"
	"os"

	"github.com/henry-l-zhou/restaurant-elo-system/internal/externals/nearbysearch"
)

func CreateNearbyPlaceData(long float64, lat float64, radius float64, place_type string) {
	params := url.Values{}
	params.Set("key", os.Getenv("API_KEY"))
	params.Set("type", place_type)

	for {
		nearby_data, err := nearbysearch.GetPlaceIDs(long, lat, radius, place_type, params)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, result := range nearby_data.Results {
			err = WriteJSONData([]interface{}{result.PlaceID, result.Name}, "place_ids", "place_id, name")

			if err != nil {
				fmt.Println(err)
			}
		}
		if nearby_data.NextPageToken != "" {
			params.Set("pagetoken", nearby_data.NextPageToken)
			continue
		}

		break
	}
}
