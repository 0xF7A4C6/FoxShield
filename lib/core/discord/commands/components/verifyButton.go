package components

import (
	"bot/lib/core/discord/verification"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func VerifyButton(s *discordgo.Session, i *discordgo.InteractionCreate) {
	T, Pass, Err := verification.CreateVerificationTask(i.Member.User.ID, i.GuildID, i.Member.User)

	if Pass {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: "You have been successfully verified!",
			},
		})

		T.EndTask(s)
		return
	}

	if Err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("%s. Please contact a staff member to resolve this issue and be verified manually.", Err),
			},
		})
		return
	}

	Err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:        fmt.Sprintf("%s Something went wrong!", utils.Emoji_warn),
					Description: "**Hey, not so fast you rascal!\nWe have detected that something is wrong, or the server is in maximum security mode.\n\nPlease complete the captcha below**",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🤖",
							},
							Label: "I am not a robot",
							Style: discordgo.LinkButton,
							URL:   fmt.Sprintf("https://proxies.gay/verify?task=%s", T.TaskId),
						},
					},
				},
			},
		},
	})

	utils.HandleError(Err)
}
