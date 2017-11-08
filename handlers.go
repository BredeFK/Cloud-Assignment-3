package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

//HandleMain main function for /
func HandleMain(w http.ResponseWriter, r *http.Request) {
	//	URL := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Dyno woken up! yai %s\n", http.StatusOK)
	/*
		switch r.Method {
		case "GET":
			HandleTest(w, r)
			//HandleGet(URL[1], w, r)
		case "POST":
			//HandlePost(w, r)
		case "DELETE":
			//HandleDelete(URL[1], w, r)
		default:
			http.Error(w, "Request not supported.", http.StatusNotImplemented)
		}*/
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "/*Here be dragons*/")
	case "POST":
		//HandlePost(w, r)
		fmt.Fprintln(w, "/*Here be dragons too*/")
		//ta inn webhook shit her.
	default:
		http.Error(w, "Request not supported.", http.StatusNotImplemented)
	}
}


func HandleAddCurrency(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.fixer.io/latest?base=EUR"
	DailyCurrencyAdder(URL)
}

func test(w http.ResponseWriter, r *http.Request){

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

	//2d data map
	var data2d map[string]map[string]float64

	//initialize the map
	data2d = make(map[string]map[string]float64)

	for i := 0; i < len(s1); i++ {
		//gets currencies from Fixer with the BASE currency
		json1, err := http.Get("http://api.fixer.io/latest?base=" + s1[i]) //+ "," + s2)
		if err != nil {
		fmt.Printf("fixer.io is not responding, %s\n", err)
		return
		}

		//data object
		var data Data

		//json decoder
		err = json.NewDecoder(json1.Body).Decode(&data)
		if err != nil { //err handler
		fmt.Printf("Error: %s\n", err)
		return
		}

		//loops through all currency codes and adds them to 2d data map
		for j := 0; j < len(s1); j++ {
			//skip identical currency codes
			if s1[i] != s1[j] {
				Add2d(data2d, s1[i], s1[j], data.Rates[s1[j]])
			}
		}
	}
	fmt.Fprintf(w, "%s\n", data2d)

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
