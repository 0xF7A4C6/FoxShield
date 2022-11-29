package events

import (
	"bot/lib/core/discord/algorithm"
	"bot/lib/utils"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	Score := algorithm.GetScore(m.User)
	fmt.Printf("[JON] %s | Score: %d | Avatar: %d, Banner: %d, Discriminator: %d, Badges: %d, Snowflake: %d\n", m.User, Score.Total, Score.Avatar, Score.Banner, Score.Discriminator, Score.Badges, Score.Snowflake)

	if Score.Total < 0 {
		Err := s.GuildBanCreate(m.GuildID, m.User.ID, 7)
		if utils.HandleError(Err) {
			return
		}

		s.ChannelMessageSend("1045677439524798637", fmt.Sprintf("`🔐` Banned `%s#%s`. Suspicious score `%d`", m.User.Username, m.User.Discriminator, Score.Total))
	}
}
