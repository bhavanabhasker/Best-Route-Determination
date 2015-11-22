package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Uber API endpoint
	Url string = "https://sandbox-api.uber.com/v1/requests/"
)

type Makerequest struct {
	Status     string `json : "status"`
	Request_id string `json : "request_id"`
	Eta        int    `json : "eta"`
}

func MakeRequest(requestid string) Makerequest {
	client := &http.Client{}
	var pm Makerequest
	param := Url
	param += requestid
	// Valid Access token generated on November 10th . Expires December 10th
	//Access Token needs to be filled here , For security reasons the access
	//token has been emailed to the Professor	
	accesstoken := ""
	authstr := fmt.Sprintf("Bearer %s", accesstoken)
	request, err := http.NewRequest("GET", param, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Authorization", authstr)
	request.Header.Add("Content-Type", "application/json")
	result, _ := client.Do(request)

	resultdata, err := ioutil.ReadAll(result.Body)
	result.Body.Close()
	if e := json.Unmarshal(resultdata, &pm); e != nil {
		log.Fatal(e)
	}

	return pm
}
