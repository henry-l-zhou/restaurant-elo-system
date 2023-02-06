package nearbysearch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createLocationRequest(long float64, lat float64, radius float64) string {
	baseURL := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?"
	location := fmt.Sprintf("location=%f,%f", long, lat)
	rad := fmt.Sprintf("radius=%f", radius)
	return baseURL + location + "&" + rad + "&"
}

func GetPlaceIDs(long float64, lat float64, radius float64, place_type string, params url.Values) (*Nearby, error) {
	// encode the request parameters and append them to the URL
	// apiURL := createLocationRequest(43.07256403835086, -89.38788040245322, 100) + params.Encode()
	apiURL := createLocationRequest(long, lat, radius) + params.Encode()
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading response body")
		return nil, err
	}

	var nearby_data Nearby
	err = json.Unmarshal(body, &nearby_data)
	if err != nil {
		fmt.Println("Error in Parsing JSON")
		return nil, err
	}

	if nearby_data.Status == "REQUEST_DENIED" {
		fmt.Println("Error in authorization to Google API")
		return nil, err
	}
	return &nearby_data, nil
}
