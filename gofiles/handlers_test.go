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

func TestGetCurrency(t *testing.T) {
	out := []string{"NOK", "EUR"}
	testValue := GetValue(out[0], out[1])
	data2d := GetCurrency()

	if testValue != data2d.Data[out[0]][out[1]] {
		t.Fatalf("ERROR expected: %v got %v", testValue, data2d.Data[out[0]][out[1]])
	}
}

/*
func TestDiscordHandler(t *testing.T) {
	var testmessage discordgo.Message
	testmessage = discordgo.Message{
		ID: "test",
		ChannelID: "371707640041963524",
		Content: "NOK to EUR",
		Type: 0,
	}
	msgbyte, _ := json.Marshal(testmessage)

	s, _ := discordgo.New("Bot Mzc3MjAwMzM1OTUwOTcwOTAw.DOJtqw.6cxZr4PpXXE6OWW_ned6mO8mizg")

	s.AddHandler(DiscordHandler)
	s.MockEvent(1, msgbyte)

	response, _ := s.ChannelMessages(testmessage.ChannelID, 1, "", "", "")

	fmt.Println("\n", response[0].Content)

}*/
