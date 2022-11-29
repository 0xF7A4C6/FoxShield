package algorithm

import "fmt"

func HandleNewJoin(GuildID string, Score *BotScore) {
	if _, ok := GuildHandlerList[GuildID]; !ok {
		GuildHandlerList[GuildID] = GuildHandler{
			GuildID:    GuildID,
			LastedJoin: map[string]BotScore{},
		}

		GuildHandlerList[GuildID].LastedJoin[Score.User.ID] = *Score
		return
	}

	Handler := GuildHandlerList[GuildID]

	if _, ok := Handler.LastedJoin[Score.User.ID]; !ok {
		Handler.LastedJoin[Score.User.ID] = *Score
	}

	if len(Handler.LastedJoin) < 5 {
		fmt.Printf("Not enought join on %s: %d/5\n", GuildID, len(Handler.LastedJoin))
		return
	}
}
