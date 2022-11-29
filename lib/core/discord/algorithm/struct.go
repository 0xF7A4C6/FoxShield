package algorithm

import "github.com/bwmarrin/discordgo"

var (
	GuildHandlerList = make(map[string]GuildHandler)
)

type BotScore struct {
	Discriminator int
	Snowflake     int
	Badges        int
	Avatar        int
	Banner        int
	Total         int
	User          *discordgo.User
}

type GuildHandler struct {
	LastedJoin map[string]BotScore
	GuildID    string
}
