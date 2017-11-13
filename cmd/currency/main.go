package main

import (
	"flag"
	"github.com/JohanAanesen/CloudTech_oblig3/gofiles"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	//starts a session with discord
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error creating Discord session,", err.Error())
		return
	}

	// Register the DiscordHandler func as a callback for Discord message events.
	// e.g whenever a message is sent in discord, the handler is triggered
	dg.AddHandler(gofiles.DiscordHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("Error opening connection,", err.Error())
		return
	}

	http.HandleFunc("/", gofiles.HandleMain)

	// Router
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
	//	http.ListenAndServe(":8080", nil)

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()

}
