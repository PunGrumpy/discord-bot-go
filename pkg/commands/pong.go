package commands

import "github.com/bwmarrin/discordgo"

type Pong struct {
	Name string
}

func (p *Pong) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
