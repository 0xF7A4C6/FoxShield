package cmd

import (
	"bot/lib/core/database"
	"bot/lib/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func SetupVerification(s *discordgo.Session, i *discordgo.InteractionCreate) {
	GuildDb := database.GetGuild(i.GuildID)
	GuildDb.VerificationRoleID = i.ApplicationCommandData().Options[0].RoleValue(s, i.GuildID).ID
	database.UpdateGuild(i.GuildID, GuildDb)

	Err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsCrossPosted,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("%s Verification required", utils.Emoji_warn),
					Description: "**We need to verify that you are not a robot.\n__Please check by clicking on the button below__ to access the other channels.**",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Label:    "Let me get into the server !",
							Style:    discordgo.DangerButton,
							CustomID: "verify_btn",
						},
					},
				},
			},
		},
	})

	utils.HandleError(Err)
}
