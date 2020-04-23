package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const YelpBuisnessSearchURL = "https://api.yelp.com/v3/businesses/search?"

func SearchBars(longitude string, latitude string, location string, limit string) (map[string]interface{}, error) {
	var res map[string]interface{}

	request, err := http.NewRequest("GET", YelpBuisnessSearchURL, nil)
	apiKey := os.Getenv("API_KEY")
	bearerKey := "Bearer " + apiKey
	request.Header.Add("Authorization", bearerKey)

	q := request.URL.Query()
	if location != "" {
		q.Add("location", location)
	}
	if longitude != "" {
		q.Add("longitude", longitude)
	}
	if latitude != "" {
		q.Add("latitude", latitude)
	}
	if limit != "" {
		q.Add("limit", limit)
	}
	q.Add("sort_by", "rating")
	q.Add("term", "bars")
	request.URL.RawQuery = q.Encode()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return res, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
