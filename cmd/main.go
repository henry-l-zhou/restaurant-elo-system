package main

import (
	"fmt"

	"github.com/henry-l-zhou/restaurant-elo-system/internal/database"
	"github.com/henry-l-zhou/restaurant-elo-system/internal/externals/nearbysearch"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Placez struct {
	place_id string
	name     string
}

func main() {
	err := godotenv.Load("../'.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	database.Connect()

	nearby_data, err := nearbysearch.GetPlaceIDs(43.072564, -89.387880, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, result := range nearby_data.Results {

		err = database.WriteJSONData([]interface{}{result.PlaceID, result.Name}, "place_ids", "place_id, name")

		if err != nil {
			fmt.Println(err)
		}

	}
}
