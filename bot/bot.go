package bot

import (
	"fmt"
	"log"
	"strings"

	"../commands"
	"../config"

	"github.com/bwmarrin/discordgo"
)

var GoBot *discordgo.Session
var BotID string

// Start is the main routine of the bot
func Start() {
	GoBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	u, err := GoBot.User("@me")

	if err != nil {
		log.Fatal(err)
	}

	BotID = u.ID
	// Handlers
	GoBot.AddHandler(messageHandler)

	// Open the connection to discord server
	err = GoBot.Open()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is starting...")
	<-make(chan struct{})
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all the messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	if m.Content[0] != '!' {
		return
	}

	command := m.Content[1:]

	args := strings.Split(command, " ")

	switch args[0] {
	case "hello":
		_, _ = s.ChannelMessageSend(m.ChannelID, commands.Hello(m.Author.Username))
		fmt.Println(m.Author.Username)
	}
}
