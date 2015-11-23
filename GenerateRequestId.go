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
	accesstoken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiJiNTJjZjYxZC00ODA0LTQwNTAtYjRmNy02Yzk4ZDMyZmQ2NjgiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjQ1OGY3MTExLWIxNDItNGE5Ny04YWQyLWQ3NDQ1MWM4YTI5ZCIsImV4cCI6MTQ0OTgyMzY2OSwiaWF0IjoxNDQ3MjMxNjY4LCJ1YWN0IjoiaHNsZEE3aXkycnRJNkxhOUNBa0tIUXIwSlo3TmllIiwibmJmIjoxNDQ3MjMxNTc4LCJhdWQiOiJ1eF9wbFNFYkhSSTMybG1XZ0tpR09KMVN6YjRWencwbyJ9.ZPHFFP8CGQEWkmY4pOwaz8pamsA6a7ePB0xWAUPfP5EKpawH8_S2V-4WWAmNVDQqtCtHUQ7afqlMD13QFPZXEAn3Ztc3L1xGY_LACHzz2ZE8we5xktVgVrfzy7fFDCnwIHAkigmzoiQnaW5DW4_Sxjtf8Fz9gwqShNYCMDGWmx-TkbP6w425CHfyBXJObjJijpv_8Z6bXp-u8bYIAkNp0OitfRPfhZpErLiFbL01R7RzdLW1BkeiHC7Dw9513xwY7GvoZBkNdVU9WU2dlfT9k-CH0vakWEqJOSp7oyTpmNmsAmqtpENH1ASUjYKK5hswfhjVdxQ91LPCJevFxb1pBA"
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
