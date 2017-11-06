package main

import (
	"net/http"
	"fmt"
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

func HandleTest(w http.ResponseWriter, r *http.Request){


}