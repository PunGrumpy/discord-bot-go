package main

import (
	"fmt"

	"github.com/PunGrumpy/discord-bot-go/pkg/bot"
	"github.com/PunGrumpy/discord-bot-go/pkg/env"
)

func main() {
	token := env.LoadVar("BOT_TOKEN")
	guildID := env.LoadVar("GUILD_ID")

	if token == "" || guildID == "" {
		fmt.Println("No BOT_TOKEN or GUILD_ID environment variable found")
		return
	}

	discordBot, err := bot.NewBot(token, guildID)
	if err != nil {
		fmt.Println("Error creating bot: ", err)
		return
	}

	go discordBot.StartServer() // start the server

	if err := discordBot.Start(); err != nil {
		fmt.Println("Error starting bot: ", err)
	}
}
