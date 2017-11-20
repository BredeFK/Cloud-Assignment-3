package gofiles

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//HandleMain main function for /
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dyno woken up! yai statuscode: %v\n", http.StatusOK)
}

// DiscordHandler is created on any channel that the autenticated bot has access to.
func DiscordHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var value float64
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	db := SetupDB()

	//sends the message from discord to dialogflow, gets back answer, base and target currency and value
	ans, base, target, amount := SendFlow(m.Content, m.Author.ID)
	//if base and target is not empty
	if base != "" || target != "" {
		value = db.GetValue(base, target)
	}

	//if amount is not empty
	if amount != "" {
		//convert string to float64
		amount2, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			fmt.Println("Somethings wrong: ", err)
			return
		} else if amount2 > 0 {
			//else return the amount times value
			//this makes it possible to ask for the rate between e.g 10 NOK in EUR
			value *= amount2
		}
	}

	//if the value is 0, do not send an answer with the value
	if value != 0 {
		//sends answer to discord with value
		s.ChannelMessageSend(m.ChannelID, ans+" "+fmt.Sprintf("%.3f", value))
	} else {
		//sends answer to discord without value
		s.ChannelMessageSend(m.ChannelID, ans)
	}

}

// SendFlow ...
func SendFlow(discMsg string, discID string) (string, string, string, string) {
	//gets api token from env var
	authToken := os.Getenv("APIAI_TOKEN")

	//testing purposes
	authToken = "5bd836a84e0747a1a091bb1a6aef9ad1"

	//url parameters
	params := url.Values{}
	params.Add("query", discMsg)
	params.Set("sessionId", discID)

	//formats the GET request to dialogflow/api.ai
	URL := fmt.Sprintf("https://api.api.ai/v1/query?V=20170712&lang=En&%s", params.Encode())
	ai, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("something wrong with the GET request to dialogflow!")
		return "", "", "", ""
	}
	//set auth with token
	ai.Header.Set("Authorization", "Bearer "+authToken)

	//send the request
	if resp, err := http.DefaultClient.Do(ai); err == nil {
		defer resp.Body.Close()

		var input APIPayload

		//decode response
		datastring, _ := ioutil.ReadAll(resp.Body)
		err := json.NewDecoder(strings.NewReader(string(datastring))).Decode(&input)
		if err != nil {
			return "", "", "", ""
		}
		//parameters from response, type interface{} sent back as type string
		//if no parameters from dialogflow, send only the answer
		if input.Result.Parameters == nil {
			return input.Result.Speech, "", "", ""
		}
		//if only some of the parameters are sent, send only the answer
		if input.Result.Parameters["baseCurrency"] == nil || input.Result.Parameters["targetCurrency"] == nil || input.Result.Parameters["number"] == nil {
			return input.Result.Speech, "", "", ""
		}
		//if the number parameter is not empty, send all data
		if input.Result.Parameters["number"] != "" {
			return input.Result.Speech, input.Result.Parameters["baseCurrency"].(string), input.Result.Parameters["targetCurrency"].(string), input.Result.Parameters["number"].(string)
		}
		//else send without number
		return input.Result.Speech, input.Result.Parameters["baseCurrency"].(string), input.Result.Parameters["targetCurrency"].(string), ""
	}
	//if it hasn't returned anything yet, return empty
	return "", "", "", ""
}
