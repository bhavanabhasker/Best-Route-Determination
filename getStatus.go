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
	accesstoken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiJiNTJjZjYxZC00ODA0LTQwNTAtYjRmNy02Yzk4ZDMyZmQ2NjgiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjQ1OGY3MTExLWIxNDItNGE5Ny04YWQyLWQ3NDQ1MWM4YTI5ZCIsImV4cCI6MTQ0OTgyMzY2OSwiaWF0IjoxNDQ3MjMxNjY4LCJ1YWN0IjoiaHNsZEE3aXkycnRJNkxhOUNBa0tIUXIwSlo3TmllIiwibmJmIjoxNDQ3MjMxNTc4LCJhdWQiOiJ1eF9wbFNFYkhSSTMybG1XZ0tpR09KMVN6YjRWencwbyJ9.ZPHFFP8CGQEWkmY4pOwaz8pamsA6a7ePB0xWAUPfP5EKpawH8_S2V-4WWAmNVDQqtCtHUQ7afqlMD13QFPZXEAn3Ztc3L1xGY_LACHzz2ZE8we5xktVgVrfzy7fFDCnwIHAkigmzoiQnaW5DW4_Sxjtf8Fz9gwqShNYCMDGWmx-TkbP6w425CHfyBXJObjJijpv_8Z6bXp-u8bYIAkNp0OitfRPfhZpErLiFbL01R7RzdLW1BkeiHC7Dw9513xwY7GvoZBkNdVU9WU2dlfT9k-CH0vakWEqJOSp7oyTpmNmsAmqtpENH1ASUjYKK5hswfhjVdxQ91LPCJevFxb1pBA"
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
