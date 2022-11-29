package events

import (
	"bot/lib/utils"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.Debug("+", fmt.Sprintf("online on %d guilds !", len(s.State.Guilds)))

	go func() {
		Status := []string{
			"/config",
			"Protects <count> servers !",
		}

		for {
			for _, S := range Status {
				s.UpdateGameStatus(0, strings.ReplaceAll(S, "<count>", fmt.Sprintf("%d", len(s.State.Guilds))))
				time.Sleep(time.Minute)
			}
		}
	}()
}
