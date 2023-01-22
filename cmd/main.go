package main

import (
	"log"
	"os"

	"github.com/Flokey82/go_discbot"
)

func main() {
	bot, err := go_discbot.NewBot("./config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := bot.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
