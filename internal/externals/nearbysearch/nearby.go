package nearbysearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func createLocationRequest(long float64, lat float64, radius float64) string {
	baseURL := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?"
	location := fmt.Sprintf("location=%f,%f", long, lat)
	rad := fmt.Sprintf("radius=%f", radius)
	return baseURL + location + "&" + rad + "&"
}

func GetPlaceIDs(long float64, lat float64, radius float64) {
	err := godotenv.Load("../../'.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	params := url.Values{}
	params.Set("key", os.Getenv("API_KEY"))

	// encode the request parameters and append them to the URL
	// apiURL := createLocationRequest(43.07256403835086, -89.38788040245322, 100) + params.Encode()
	apiURL := createLocationRequest(long, lat, radius) + params.Encode()
	fmt.Println(apiURL)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading response body")
		return
	}

	var nearby_data Nearby
	err = json.Unmarshal(body, &nearby_data)
	if err != nil {
		fmt.Println("Error in Parsing JSON")
	}

	fmt.Println(nearby_data)
}
