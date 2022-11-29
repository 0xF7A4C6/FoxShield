package events

import (
	"bot/lib/core/database"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	database.RemoveGuild(g.ID)
	utils.Debug("+", fmt.Sprintf("Leaved guild %s", g.ID))
}
