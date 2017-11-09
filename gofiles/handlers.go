package gofiles

import (
	"fmt"
	"net/http"
)

//HandleMain main function for /
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dyno woken up! yai statuscode: %v\n", http.StatusOK)
}
