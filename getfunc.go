package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetCurrency gets the currency from string URL
func GetCurrency() Data2d {

	//Currency codes
	s1 := []string{
		"AUD", "BGN", "BRL", "CAD",
		"CHF", "CNY", "CZK", "DKK",
		"EUR", "GBP", "HKD", "HRK",
		"HUF", "IDR", "ILS", "INR",
		"JPY", "KRW", "MXN", "MYR",
		"NOK", "NZD", "PHP", "PLN",
		"RON", "RUB", "SEK", "SGD",
		"THB", "TRY", "USD", "ZAR"}

	//initialize the map
	var data2d Data2d
	data2d.Data = make(map[string]map[string]float64)

	//sets date
	data2d.Date = time.Now().Format("2006-01-02")

	for i := 0; i < len(s1); i++ {
		//gets currencies from Fixer with the BASE currency
		json1, err := http.Get("http://api.fixer.io/latest?base=" + s1[i]) //+ "," + s2)
		if err != nil {
			fmt.Printf("fixer.io is not responding, %s\n", err)
			panic(err)
		}

		//data object
		var data Data

		//json decoder
		err = json.NewDecoder(json1.Body).Decode(&data)
		if err != nil { //err handler
			fmt.Printf("Error: %s\n", err)
			panic(err)
		}

		//loops through all currency codes and adds them to 2d data map
		for j := 0; j < len(s1); j++ {
			//skip identical currency codes
			if s1[i] != s1[j] {
				//add data to 2d map
				Add2d(data2d.Data, s1[i], s1[j], data.Rates[s1[j]])
			}
		}
	}

	return data2d
}

//Add2d adds base, target and value of a currency to the 2d map
func Add2d(m map[string]map[string]float64, base string, target string, value float64) {
	mm, ok := m[base]
	//initializes the child maps
	if !ok {
		mm = make(map[string]float64)
		m[base] = mm
	}
	mm[target] = value
}

// GetValue gets value from db
func GetValue(s1 string, s2 string) float64 {

	// Get today's date in date format
	tempToday := time.Now().Local()
	today := tempToday.Format("2006-01-02")

	// Set up the database
	db := SetupDB()
	db.Init()

	// Get today's currencies for today
	data2d, ok := db.GetLatest(today)

	// If there isn't any data in the db for today
	if ok == false {

		// Try to get data from yesterday
		tempToday = time.Now().Local().AddDate(0, 0, -1)
		yesterday := tempToday.Format("2006-01-02")
		data2d, ok = db.GetLatest(yesterday)

		// If there's still not any data: log error to heroku
		if ok == false {
			log.Println("Could not get any data from today or yesterday:/", 404)
		}
	}
	return data2d.Data[s1][s2]
}
