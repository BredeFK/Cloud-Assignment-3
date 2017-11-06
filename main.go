package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net/http"
)

func main() {

	//token := os.Getenv("TOKEN")
	//dg, err := discordgo.New(token)
	dg, err := discordgo.New("Mzc3MjAwMzM1OTUwOTcwOTAw.DOJtqw.6cxZr4PpXXE6OWW_ned6mO8mizg")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}


	/*http.HandleFunc("/", HandleMain)


	http.ListenAndServe(":8080", nil)*/
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
//	if m.Author.ID == s.State.User.ID {
//		return
//	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "fuck" {
		s.ChannelMessageSend(m.ChannelID, "you!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}