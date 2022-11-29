package commands

import (
	"bot/lib/core/discord/commands/cmd"
	"github.com/bwmarrin/discordgo"
)

var (
	CommandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"setup":  cmd.SetupVerification,
		"config": cmd.ConfigServer,
	}
)
