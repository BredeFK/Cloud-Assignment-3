package gofiles

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(HandleMain).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}
}

func TestSendFlow(t *testing.T) {
	testMessage := "10 NOK TO EUR"
	in := []string{"NOK", "EUR", "10"}

	_, base, target, amount := SendFlow(testMessage, "12412413")

	if base != in[0] {
		t.Fatalf("ERROR: Expected %s got %s", in[0], base)
	}
	if target != in[1] {
		t.Fatalf("ERROR: Expected %s got %s", in[1], target)
	}
	if amount != in[2] {
		t.Fatalf("ERROR: Expected %s got %s", in[2], amount)
	}
}
