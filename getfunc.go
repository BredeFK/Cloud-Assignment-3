package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetCurrency gets the currency from string URL
func GetCurrency(URL string) Currency{
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("User-Agent", "Assignment")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr.Error())
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr.Error())
	}

	currency := Currency{}
	jsonErr := json.Unmarshal(body, &currency)
	if jsonErr != nil {
		log.Fatal(jsonErr.Error())
	}

	return currency
}