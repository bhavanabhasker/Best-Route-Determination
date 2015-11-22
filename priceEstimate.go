package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	// Uber API endpoint
	APIUrl string = "https://sandbox-api.uber.com/v1/%s%s"
)

// Uber price estimate
type PriceEstimate struct {
	ProductId       string  `json:"product_id"`
	CurrencyCode    string  `json:"currency_code"`
	DisplayName     string  `json:"display_name"`
	Estimate        string  `json:"estimate"`
	LowEstimate     int     `json:"low_estimate"`
	HighEstimate    int     `json:"high_estimate"`
	SurgeMultiplier float64 `json:"surge_multiplier"`
	Duration        int     `json:"duration"`
	Distance        float64 `json:"distance"`
}

// Send HTTP request to Uber API
func getResponse(start_latitude float64, start_longitude float64, end_latitude float64,
	end_longitude float64) []byte {

	endpoint := "estimates/price"

	params := map[string]string{
		"start_latitude":  strconv.FormatFloat(start_latitude, 'f', 2, 32),
		"start_longitude": strconv.FormatFloat(start_longitude, 'f', 2, 32),
		"end_latitude":    strconv.FormatFloat(end_latitude, 'f', 2, 32),
		"end_longitude":   strconv.FormatFloat(end_longitude, 'f', 2, 32),
	}
	urlParams := "?"
	params["server_token"] = "QCDOAFBQJRehdFg-LYijrgDQWazE7AP-eNyiqVjA"
	for k, v := range params {
		if len(urlParams) > 1 {
			urlParams += "&"
		}
		urlParams += fmt.Sprintf("%s=%s", k, v)

	}

	url := fmt.Sprintf(APIUrl, endpoint, urlParams)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return data
}
