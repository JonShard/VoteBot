package main

import (
	"Votebot/votebot/bot"
	"Votebot/votebot/cfg"
	"Votebot/votebot/db"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg.ReadConfigFile()
	bot.Init()
	db.Init()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Cxt.Session.Close()
	cfg.WriteConfigFile()
}
