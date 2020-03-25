package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/kelseyhightower/envconfig"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Token string
}

func messageHandle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		fmt.Printf("Message from bot\n")
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	} else if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

func main() {
	fmt.Println("vim-go")
	var c Config
	err := envconfig.Process("fcbot", &c)
	if err != nil {
		fmt.Errorf("ERROR: processing environment: %v", err)
	}

	spew.Dump(c)
	bot, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		fmt.Errorf("ERROR: could not create discord session!: %v", err)
	}

	bot.AddHandler(messageHandle)

	err = bot.Open()
	if err != nil {
		fmt.Errorf("ERROR: could not open bot: %v", err)
	}
	fmt.Printf("INFO: started!\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	bot.Close()
}
