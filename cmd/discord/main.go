package main

import (
	"fmt"

	"github.com/PunGrumpy/discord-bot-go/pkg/bot"
	"github.com/PunGrumpy/discord-bot-go/pkg/env"
)

func main() {
	token := env.LoadVar("BOT_TOKEN")

	if token == "" {
		fmt.Println("No BOT_TOKEN environment variable found")
		return
	}

	discordBot, err := bot.NewBot(token)
	if err != nil {
		fmt.Println("Error creating bot: ", err)
		return
	}

	if err := discordBot.Start(); err != nil {
		fmt.Println("Error starting bot: ", err)
	}
}
