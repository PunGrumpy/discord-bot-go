package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}
