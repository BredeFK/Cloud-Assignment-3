package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JohanAanesen/CloudTech_oblig3/gofiles"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Variables used for command line parameters
var (
	Token string
)

// init
func init() {

	flag.StringVar(&Token, "t", os.Getenv("DISCORD_TOKEN"), "Bot Token")
	flag.Parse()
}

func main() {

	//token := os.Getenv("TOKEN")
	//dg, err := discordgo.New(token)
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	http.HandleFunc("/", gofiles.HandleMain)
	http.HandleFunc("/webhook", gofiles.HandleWebhook)

	//Router
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
	//	http.ListenAndServe(":8080", nil)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}

// messageCreate is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	ans, base, target := SendFlow(m.Content, m.Author.ID)

	value := gofiles.GetValue(base, target)

	if value != 0 {
		s.ChannelMessageSend(m.ChannelID, ans+" "+fmt.Sprint(value))
	}
}

// SendFlow ...
func SendFlow(discMsg string, discID string) (string, string, string) {
	authToken := os.Getenv("APIAI_TOKEN")

	params := url.Values{}
	params.Add("query", discMsg)
	params.Set("sessionId", discID)

	URL := fmt.Sprintf("https://api.api.ai/v1/query?V=20170712&lang=En&%s", params.Encode())
	ai, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("something wrong with the GET request to dialogflow!")
		return "", "", ""
	}

	ai.Header.Set("Authorization", "Bearer "+authToken)

	if resp, err := http.DefaultClient.Do(ai); err != nil {
		return "", "", ""
	} else {
		defer resp.Body.Close()

		var input gofiles.ApiPayload
		datastring, _ := ioutil.ReadAll(resp.Body)
		err := json.NewDecoder(strings.NewReader(string(datastring))).Decode(&input)
		if err != nil {
			return "", "", ""
		}

		return input.Result.Speech, input.Result.Parameters["baseCurrency"], input.Result.Parameters["targetCurrency"]

	}
}
