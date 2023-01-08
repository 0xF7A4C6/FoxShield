package events

import (
    "bot/lib/core/discord/algorithm"
    "bot/lib/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "`🏓` Pong!",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  fmt.Sprintf("%s Bot", utils.Emoji_timeout),
					Value: fmt.Sprintf("`%dms`", time.Now().Sub(m.Timestamp).Milliseconds()),
				},
				{
					Name:  fmt.Sprintf("%s Websocket", utils.Emoji_timeout),
					Value: fmt.Sprintf("`%dms`", s.HeartbeatLatency().Milliseconds()),
				},
			},
		})
	}

    go algorithm.HandleNewJoin(s, m.GuildID, algorithm.GetScore(m.Author))
}
