package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	client "github.com/mgiles717/moxbot/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var bot_token = os.Getenv("BOT_TOKEN")

	client.InitialiseEconomy()
	client.StartBot(bot_token)
}
