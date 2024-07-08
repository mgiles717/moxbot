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

	// If data/economy.parquet does not exist, run client.InitialiseEconomy() to create it

	if _, err := os.Stat("data/economy.parquet"); os.IsNotExist(err) {
		client.InitialiseEconomy()
	}
	client.StartBot(bot_token)
}
