package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	defaultMemberPermissions int64 = discordgo.PermissionAdministrator
	dmPermission                   = false

	Commands = []*discordgo.ApplicationCommand{
		{
			Name:                     "setup",
			Description:              "Setup verification in the current channel",
			Type:                     discordgo.ChatApplicationCommand,
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "verification_role_id",
					Description:  "Role to add after being verified.",
					Type:         discordgo.ApplicationCommandOptionRole,
					Required:     true,
					Autocomplete: false,
				},
			},
		},
		{
			Name:                     "config",
			Description:              "Edit server configuration.",
			Type:                     discordgo.ChatApplicationCommand,
			DefaultMemberPermissions: &defaultMemberPermissions,
			DMPermission:             &dmPermission,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "verification_role_id",
					Description:  "Role to add after being verified.",
					Type:         discordgo.ApplicationCommandOptionRole,
					Required:     false,
					Autocomplete: false,
				},
				{
					Name:         "verification_level",
					Description:  "Set captcha verification level.",
					Type:         discordgo.ApplicationCommandOptionInteger,
					Required:     false,
					Autocomplete: false,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "VerificationLevelShowNever",
							Value: 0,
						},
						{
							Name:  "VerificationLevelShowAlways",
							Value: 1,
						},
						{
							Name:  "VerificationLevelShowSuspicious",
							Value: 2,
						},
						{
							Name:  "VerificationLevelShowSuspiciousAndBlockBot",
							Value: 3,
						},
					},
				},
				{
					Name:         "verification_log_channel_id",
					Description:  "Channel where logs will be send when user pass the verification process.",
					Type:         discordgo.ApplicationCommandOptionChannel,
					Required:     false,
					Autocomplete: false,
				},
			},
		},
	}
)
