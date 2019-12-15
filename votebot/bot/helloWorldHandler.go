package bot

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// validates the message and splits it into an array of substings split on space.
func parseMessage(m *discordgo.MessageCreate) ([]string, error) {

	if len(m.Content) < 5 {
		return nil, errors.New("message too short. No commands exist shorter than 5 chars")
	}

	substrings := strings.Split(m.Content, " ")

	return substrings, nil
}

// HelloWorldHandler responds to "!Hello" with "World!" in a discord text channel.
func HelloWorldHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

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
