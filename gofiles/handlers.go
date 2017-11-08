package gofiles

import (
	"fmt"
	"net/http"
)

//HandleMain main function for /
func HandleMain(w http.ResponseWriter, r *http.Request) {
	//	URL := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "Dyno woken up! yai %v\n", http.StatusOK)
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

// HandleWebhook handles webhooks
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

// HandleAddCurrency adds currencies to db
func HandleAddCurrency(w http.ResponseWriter, r *http.Request) {
	DailyCurrencyAdder()
}
