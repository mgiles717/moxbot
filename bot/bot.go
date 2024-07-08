package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)


const prefix = "!"


func checkNilErr(e error) {
	if e != nil {
		log.Fatal("Error message")
	}
}


func StartBot(bot_token string) {
	discord, err := discordgo.New("Bot " + bot_token)
	checkNilErr(err)
	
	discord.AddHandler(NewMessage)
	discord.AddHandler(NewMessage)

	discord.Open()
	defer discord.Close()

	fmt.Println("Starting bot...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}


func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, prefix+"hello"):
		discord.ChannelMessageSend(message.ChannelID, "Hello!")
	case strings.Contains(message.Content, prefix+"balance"):
		CheckBalance(message.Author.ID)
		discord.ChannelMessageSend(message.ChannelID, "You have 0 moxcoins!")
	}
}

