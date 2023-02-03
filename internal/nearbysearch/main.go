package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func CreateLocationRequest(long float64, lat float64, radius float64) string {
	baseURL := "https://maps.googleapis.com/maps/api/place/nearbysearch/json?"
	location := fmt.Sprintf("location=%f,%f", lat, long)
	rad := fmt.Sprintf("radius=%f", radius)
	return baseURL + location + "&" + rad
}

func main() {
	err := godotenv.Load("../../'.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	params := url.Values{}
	params.Set("api_key", os.Getenv("API_KEY"))

	// encode the request parameters and append them to the URL
	apiURL := CreateLocationRequest(43.07256403835086, -89.38788040245322, 100) + params.Encode()

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
