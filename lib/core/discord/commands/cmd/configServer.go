package cmd

import (
	"bot/lib/core/database"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ConfigServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	GuildDb := database.GetGuild(i.GuildID)
	Data := i.ApplicationCommandData()

	for _, Option := range Data.Options {
		switch Option.Name {
		case "verification_role_id":
			GuildDb.VerificationRoleID = Option.RoleValue(s, i.GuildID).ID

		case "verification_level":
			GuildDb.VerificationLevel = int(Option.IntValue())

		case "verification_log_channel_id":
			GuildDb.VerificationLogChannelID = Option.ChannelValue(s).ID
		}
	}

	database.UpdateGuild(i.GuildID, GuildDb)

	Err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsCrossPosted,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("%s Configuration saved.", utils.Emoji_green),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:  fmt.Sprintf("%s verification_role_id", utils.Emoji_profil),
							Value: fmt.Sprintf("`%s`", GuildDb.VerificationRoleID),
						},
						{
							Name:  fmt.Sprintf("%s verification_level", utils.Emoji_profil),
							Value: fmt.Sprintf("`%d`", GuildDb.VerificationLevel),
						},
						{
							Name:  fmt.Sprintf("%s verification_log_channel_id", utils.Emoji_profil),
							Value: fmt.Sprintf("`%s`", GuildDb.VerificationLogChannelID),
						},
					},
				},
			},
		},
	})

	utils.HandleError(Err)
}
