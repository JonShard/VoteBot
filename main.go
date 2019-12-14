package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

// Token contains the authorization token for communicating with discord.
type config struct {
	BotToken string `yaml:"botToken"`
}

// validates the message and splits it into an array of substings split on space.
func parseMessage(m *discordgo.MessageCreate) ([]string, error) {

	if len(m.Content) < 5 {
		return nil, errors.New("message too short. No commands exist shorter than 5 chars")
	}

	substrings := strings.Split(m.Content, " ")

	return substrings, nil
}

// reads config file from file into struct 'config'.
func readConfigFile(c *config) *config {

	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func hellowWorldHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	message, err := parseMessage(m)
	if err != nil {
		fmt.Print(err)
		return
	}

	if message[0] == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "World!")
	}
}

func main() {

	// Load environment file:
	cfg := config{}
	readConfigFile(&cfg)

	fmt.Println("Loaded config file: \nToken: " + cfg.BotToken)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(hellowWorldHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
