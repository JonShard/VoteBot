package main

import (
	"Votebot/votebot/bot"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

// Token contains the authorization token for communicating with discord.
type config struct {
	BotToken        string `yaml:"botToken"`
	VotesPerUser    int    `yaml:"votesPerUser"`
	VotesPerPateron int    `yaml:"votesPerPatreon"`
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
	dg.AddHandler(bot.HelloWorldHandler)

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
