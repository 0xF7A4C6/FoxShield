package discord

import (
	"bot/lib/core/discord/events"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	BotSession *discordgo.Session
)

func InitEvents(Dg *discordgo.Session) {
	Dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	Dg.AddHandler(events.InteractionCreate)
	utils.Debug("+", "InteractionCreate")

	Dg.AddHandler(events.GuildMemberAdd)
	utils.Debug("+", "GuildMemberAdd")

	Dg.AddHandler(events.MessageCreate)
	utils.Debug("+", "MessageCreate")

	Dg.AddHandler(events.GuildCreate)
	utils.Debug("+", "GuildCreate")

	Dg.AddHandler(events.GuildDelete)
	utils.Debug("+", "GuildDelete")

	Dg.AddHandler(events.Ready)
	utils.Debug("+", "Ready")
}

func InitBot(Token string) error {
	Dg, Err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if utils.HandleError(Err) {
		return Err
	}

	InitEvents(Dg)

	Err = Dg.Open()
	if utils.HandleError(Err) {
		return Err
	}

	BotSession = Dg
	return nil
}
