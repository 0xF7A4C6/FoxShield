package events

import (
	"bot/lib/core/database"
	"bot/lib/core/discord/commands"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	_, Err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, g.ID, commands.Commands)
	if Err != nil {
		panic(Err)
	}
	utils.Debug("+", fmt.Sprintf("ApplicationCommandBulkOverwrite: %s", g.ID))

	database.AddGuild(g.ID)
}
