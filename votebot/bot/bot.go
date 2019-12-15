package bot

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const helpText = "Votebot manages voting on what songs are available. Commands:\n !help Prints this help text.\nMORE"
const sorryText = "Sorry, something went wrong there."

// Context holds the neccesary information to communicate with hte server.
type Context struct {
	StartTime time.Time
	Session   *discordgo.Session
}

// Cxt holds the neccesary information to communicate with hte server.
var Cxt Context

// Init initiallizes the bot, by creating a session and registering the handler.
func Init() {

	Cxt = Context{}
	Cxt.StartTime = time.Now()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Cfg.BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	Cxt.Session = dg
	Cxt.Session.AddHandler(RouterHandler)

	// Open a websocket connection to Discord and begin listening.
	err = Cxt.Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

// validates the message and splits it into an array of substings split on space.
func parseMessage(m *discordgo.MessageCreate) ([]string, error) {

	if len(m.Content) < 5 {
		return nil, errors.New("message too short. No commands exist shorter than 5 chars")
	}
	substrings := strings.Split(m.Content, " ")

	return substrings, nil
}

// RouterHandler is responsible for parsing the users command and rounding it to the correct function.
func RouterHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	command, err := parseMessage(m)
	if err != nil {
		fmt.Printf("Parsing a message threw an error: err %v", err)
		s.ChannelMessageSend(m.ChannelID, sorryText)
		return
	}

	// SetChannel is the only command that is not spesific to the slected channel. Special case.
	if m.ChannelID != Cfg.ChannelID {
		if command[0] == "!setChannel" {
			err = SetChannel(command, s, m)
		}
		return
	}

	switch strings.ToLower(command[0]) {
	case "!hello":
		s.ChannelMessageSend(m.ChannelID, "World!")
		break
	case "!help":
		s.ChannelMessageSend(m.ChannelID, helpText)
		break

	case "!displaylist":
		break

	case "!vote":
		break

	case "!search":
		break

	case "!openvotes":
		break

	case "!closevotes":
		break

	case "!setsonglimit":
		break

	case "!setvotecount":
		break

	case "!setpateronvotecount":
		break
	}

	if err != nil {
		fmt.Printf("A command threw an error: err %v", err)
		s.ChannelMessageSend(m.ChannelID, sorryText)
	}
}

// SetChannel changes the channelID the bot looks for messages in.
func SetChannel(command []string, s *discordgo.Session, m *discordgo.MessageCreate) error {

	channel, _ := Cxt.Session.Channel(m.ChannelID) // We are guarantied the channel exist.
	guildID := channel.GuildID
	if guildID == "" {
		fmt.Println("User tried to use setChannel in DMs.")
		return errors.New("can not set a text channel that is not within a guild")
	}
	Cfg.ChannelID = m.ChannelID
	Cxt.Session.ChannelMessageSend(m.ChannelID, "New text channel set to "+channel.Name)
	fmt.Println("New text channel set to \"" + channel.Name + "\" in guild with ID: " + guildID)
	return nil
}
