package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"os"
	"log"
)

type NeutrinoCredentials struct {
	UserId string
	ApiKey string
}

func main() {
	fd, err := os.Open("credentials.json")
	if err != nil {
		log.Fatal("unable to open credentials.json")
	}

	var apiCreds NeutrinoCredentials
	err = json.NewDecoder(fd).Decode(&apiCreds)
	if err != nil {
		log.Fatal("cannot parse json in NeutrinoCredentials")
	}

	params := url.Values{}
	params.Set("user-id", apiCreds.UserId)
	params.Set("api-key", apiCreds.ApiKey)
	params.Set("ip", "162.209.104.195")

	resp, err := http.PostForm("https://neutrinoapi.com/ip-info", params)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var result map[string]interface{}
			err := json.Unmarshal(body, &result)
			if err == nil {
				fmt.Println(result["country"])
				fmt.Println(result["country-code"])
				fmt.Println(result["region"])
				fmt.Println(result["city"])
			}
		}
	}
}