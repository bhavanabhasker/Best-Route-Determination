package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	// Uber API endpoint
	UberURL string = "https://sandbox-api.uber.com/v1/requests"
)

// Uber ride request
type RideRequest struct {
	RequestId string `json:"request_id"`
	Status    string `json:"status"`
	Eta       int    `json:"eta"`
}
type Request struct {
	product_id      string
	start_latitude  float64
	start_longitude float64
	end_latitude    float64
	end_longitude   float64
}

func GenerateRequestId(request Request) RideRequest {
	var pe RideRequest
	client := &http.Client{}
	// Valid Access token generated on November 10th . Expires December 10th
	//Access Token needs to be filled here , For security reasons the access
	//token has been emailed to the Professor
	accesstoken := ""
	authstr := fmt.Sprintf("Bearer %s", accesstoken)
	payload, _ := json.Marshal(map[string]string{
		"product_id":      request.product_id,
		"start_latitude":  strconv.FormatFloat(request.start_latitude, 'f', 2, 32),
		"start_longitude": strconv.FormatFloat(request.start_longitude, 'f', 2, 32),
		"end_latitude":    strconv.FormatFloat(request.end_latitude, 'f', 2, 32),
		"end_longitude":   strconv.FormatFloat(request.end_longitude, 'f', 2, 32),
	})

	req, err := http.NewRequest("POST", UberURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", authstr)
	req.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(req)

	data, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if e := json.Unmarshal(data, &pe); e != nil {
		log.Fatal(e)
	}

	return pe
}
