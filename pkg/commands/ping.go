package commands

import "github.com/bwmarrin/discordgo"

type Ping struct {
	Name string
}

func (p *Ping) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}
