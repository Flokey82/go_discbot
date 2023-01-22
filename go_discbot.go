// Package go_discbot provides a simple wrapper for the Discord API
package go_discbot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot represents the discord bot.
type Bot struct {
	BotID  string
	Config *Config
	GoBot  *discordgo.Session
}

// NewBot creates a new Bot struct using the given configuration file.
func NewBot(configFile string) (*Bot, error) {
	config, err := ReadConfig(configFile)
	if err != nil {
		return nil, err
	}
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}

	u, err := goBot.User("@me")
	if err != nil {
		return nil, err
	}

	return &Bot{
		BotID:  u.ID,
		Config: config,
		GoBot:  goBot,
	}, nil
}

// Run starts the bot and waits for a signal to stop.
func (b *Bot) Run() error {
	if err := b.Start(); err != nil {
		return err
	}

	defer b.Stop()

	<-make(chan struct{})
	return nil
}

// Start starts the bot.
func (b *Bot) Start() error {
	b.GoBot.AddHandler(b.messageHandler)
	if err := b.GoBot.Open(); err != nil {
		return err
	}
	fmt.Println("Bot is running !")
	return nil
}

// Stop stops the bot and closes the connection.
func (b *Bot) Stop() error {
	return b.GoBot.Close()
}

const (
	CmdPing = "ping"
	CmdTell = "tell"
)

// messageHandler handles the messages sent to the bot.
func (b *Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == b.BotID || !strings.HasPrefix(m.Content, b.Config.BotPrefix) {
		return
	}

	cmd := strings.TrimPrefix(m.Content, b.Config.BotPrefix)
	switch {
	case strings.HasPrefix(cmd, CmdPing):
		if _, err := s.ChannelMessageSend(m.ChannelID, "pong"); err != nil {
			fmt.Println(err.Error())
			return
		}
	case strings.HasPrefix(cmd, CmdTell):
		if _, err := s.ChannelMessageSend(m.ChannelID, "I'm a bot"); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
