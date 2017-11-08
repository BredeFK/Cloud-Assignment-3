package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// GetCurrency gets the currency from string URL

func GetCurrency() Data2d{

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
	data2d = make(map[string]map[string]float64)

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
				Add2d(data2d, s1[i], s1[j], data.Rates[s1[j]])
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