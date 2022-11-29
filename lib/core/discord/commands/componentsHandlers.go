package commands

import (
	"bot/lib/core/discord/commands/components"
	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"verify_btn": components.VerifyButton,
	}
)
