package algorithm

import (
	"bot/lib/utils"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func HandleNewJoin(s *discordgo.Session, GuildID string, Score *BotScore) {
	if _, ok := GuildHandlerList[GuildID]; !ok {
		GuildHandlerList[GuildID] = GuildHandler{
			GuildID:       GuildID,
			LastedJoin:    map[string]BotScore{},
			BannedAccount: 0,
		}

		GuildHandlerList[GuildID].LastedJoin[Score.User.ID] = *Score
		return
	}

	Handler := GuildHandlerList[GuildID]
	if _, ok := Handler.LastedJoin[Score.User.ID]; !ok {
		Handler.LastedJoin[Score.User.ID] = *Score
	}

	LastedJoins := make([]BotScore, 0, len(Handler.LastedJoin))

	for _, value := range Handler.LastedJoin {
		if value.Total < 50 {
			LastedJoins = append(LastedJoins, value)
		}
	}

	if len(LastedJoins) < 2 {
		fmt.Printf("Not enought join on %s: %d/2\n", GuildID, len(LastedJoins))
		return
	}

	// check if server is in raid
	var Max int
	if len(LastedJoins) > 5 {
		Max = 5
	} else {
		Max = len(LastedJoins)
	}

	FirstJoinSecond := time.Now().Sub(LastedJoins[:Max][0].JoinTime).Seconds()

	for _, Member := range LastedJoins[:Max] {
		Ban := false

		if FirstJoinSecond < 15 && Member.Total <= 0 {
			Ban = true
		}

		if Ban {
			delete(Handler.LastedJoin, Member.User.ID)

			Err := s.GuildBanCreate(GuildID, Member.User.ID, 7)
			if utils.HandleError(Err) {
				continue
			}

			Handler.BannedAccount++
			fmt.Println(Handler.BannedAccount)
			go s.ChannelMessageSend("1047042742796169218", fmt.Sprintf("`🔐` Banned `%s#%s`. Suspicious score `%d`", Member.User.Username, Member.User.Discriminator, Score.Total))
		}
	}
}
