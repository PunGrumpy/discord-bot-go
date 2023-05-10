package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PunGrumpy/discord-bot-go/pkg/commands"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Token      string
	Connection *discordgo.Session
	Commands   []commands.Command
}

func NewBot(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	discordBot := &Bot{
		Token:      token,
		Connection: dg,
		Commands: []commands.Command{
			&commands.Ping{},
			&commands.Pong{},
		},
	}

	discordBot.Connection.AddHandler(discordBot.messageCreate)

	return discordBot, nil
}

func (discordBot *Bot) Start() error {
	err := discordBot.Connection.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return err
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discordBot.Connection.Close()

	return nil
}

func (discordBot *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	for _, command := range discordBot.Commands {
		command.Execute(s, m)
	}
}
