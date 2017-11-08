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

func test(w http.ResponseWriter, r *http.Request){

	s1 := []string{"EUR", "NOK", "USD", "JPY"}
	type data2d map[string]Data

	var shit data2d

	for i := 0; i < len(s1); i++ {
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

		shit[s1[i]] = data

	}
	fmt.Println(shit)
}
