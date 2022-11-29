package events

import (
	"bot/lib/core/discord/commands"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {

	case discordgo.InteractionApplicationCommand:
		if h, ok := commands.CommandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}

	case discordgo.InteractionMessageComponent:
		if h, ok := commands.ComponentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}

	default:
		utils.Debug("-", fmt.Sprintf("Interation not catched: %s", i.Type))
	}
}
