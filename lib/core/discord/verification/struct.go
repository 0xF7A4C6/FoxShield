package verification

import (
	"bot/lib/core/discord/algorithm"
	"time"
)

var (
	BehaviorSuspicious = 0
	BehaviorNormal     = 1
	BehaviorSpam       = 2
	BehaviorSafe       = 3
)

type Challenge struct {
	InToken     string
	OutResponse string
}

type Task struct {
	Score         *algorithm.BotScore
	Passed        bool
	MemberID      string
	GuildID       string
	TaskId        string
	StartTime     time.Time
	Level         string
	Behavior      int
	TaskChallenge Challenge
}
