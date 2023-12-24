package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type RailwayAPIClient struct {
	api_key  string
	api_host string
	base_url string
}

var (
	railwayClient     *RailwayAPIClient
	railwayClientOnce sync.Once
)

func GetRailwayClient() *RailwayAPIClient {
	railwayClientOnce.Do(func() {
		initializeRailwayClient()
	})

	return railwayClient
}

func initializeRailwayClient() {
	api_key := os.Getenv("X_RapidAPI_Key")
	api_host := os.Getenv("git X_RapidAPI_Host")
	railwayClient = &RailwayAPIClient{
		api_key:  api_key,
		api_host: api_host,
		base_url: "https://irctc1.p.rapidapi.com/api",
	}
}

func (c *RailwayAPIClient) GetTrains(fromStationCode string, toStationCode string, date string) (map[string]interface{}, error) {
	endpoint := c.base_url + "/v3/trainBetweenStations"
	url, _ := url.Parse(endpoint)
	payload := url.Query()
	payload.Set("dateOfJourney", date)
	payload.Set("fromStationCode", fromStationCode)
	payload.Set("toStationCode", toStationCode)
	url.RawQuery = payload.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)

	req.Header.Add("X-RapidAPI-Key", c.api_key)
	req.Header.Add("X-RapidAPI-Host", c.api_host)
	res, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error:", err)
	}
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	defer res.Body.Close()
	return responseData, err
}
