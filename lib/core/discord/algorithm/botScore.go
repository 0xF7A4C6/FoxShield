package algorithm

import (
	"bot/lib/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (S *BotScore) CheckAvatar() {
	if strings.Contains(S.User.AvatarURL("32"), "gif") {
		S.Avatar = 10
	} else if strings.Contains(S.User.AvatarURL("32"), "png") {
		S.Avatar = 1
	} else {
		S.Avatar = -2
	}
}

func (S *BotScore) CheckBanner() {
	if strings.Contains(S.User.BannerURL("32"), "gif") {
		S.Banner = 10
	} else if strings.Contains(S.User.BannerURL("32"), "png") {
		S.Banner = 1
	} else {
		S.Banner = -2
	}
}

func (S *BotScore) CheckDiscriminator() {
	for _, Num := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		var Dups int

		for _, Discrim := range S.User.Discriminator {
			if string(Discrim) == Num {
				Dups++
			}
		}

		if Dups != 1 {
			S.Discriminator += Dups
		}
	}
}

func (S *BotScore) CheckFlags() {
	/*
		class UserFlags(Enum):
		staff = 1
		partner = 2
		hypesquad = 4
		bug_hunter = 8
		mfa_sms = 16
		premium_promo_dismissed = 32
		hypesquad_bravery = 64
		hypesquad_brilliance = 128
		hypesquad_balance = 256
		early_supporter = 512
		team_user = 1024
		system = 4096
		has_unread_urgent_messages = 8192
		bug_hunter_level_2 = 16384
		verified_bot = 65536
		verified_bot_developer = 131072
		discord_certified_moderator = 262144
		bot_http_interactions = 524288
		spammer = 1048576
		active_developer = 4194304
	*/

	switch S.User.PublicFlags {
	case 1, 2, 4, 8, 512, 1024, 4096, 16384, 131072, 65536, 262144:
		S.Badges += 100

	// spammer
	case 1048576:
		S.Badges -= 50

	// hypesquad
	case 64, 128, 256:
		S.Badges += 25

	// active_developer
	case 4194304:
		S.Badges += 30
	}
}

func (S *BotScore) CheckSnowflake() {
	ID, _ := strconv.Atoi(S.User.ID)
	Timestamp := fmt.Sprintf("%d", ((ID>>22)+1420070400000)/1000)

	I, Err := strconv.ParseInt(Timestamp, 10, 64)
	if utils.HandleError(Err) {
		S.Snowflake -= 2
	}

	Diff := time.Now().Sub(time.Unix(I, 0)).Hours()

	//fmt.Println(Diff,Diff / 24,Diff / (365.24*24))

	if len(S.User.ID) != 18 {
		S.Snowflake += 15
	}

	if Diff < 1 {
		S.Snowflake -= 100
	}

	if (Diff / 24) < 1 {
		S.Snowflake -= 25
	}

	if (Diff / 24) < 3 {
		S.Snowflake -= 15
	}

	if (Diff / 24) < 5 {
		S.Snowflake -= 5
	}

	if (Diff / 24) >= 25 {
		S.Snowflake += 3
	}

	if (Diff / (365.24 * 24)) >= 5 {
		S.Snowflake += 4
	}

	if (Diff / (365.24 * 24)) >= 4 {
		S.Snowflake += 2
	}

	if (Diff / (365.24 * 24)) >= 3 {
		S.Snowflake += 25
	}

	if (Diff / (365.24 * 24)) >= 2 {
		S.Snowflake += 15
	}

	if (Diff / (365.24 * 24)) >= 1 {
		S.Snowflake += 10
	}
}

func GetScore(User *discordgo.User) *BotScore {
	Score := &BotScore{
		User: User,
	}

	Score.CheckAvatar()
	//Score.CheckBanner()
	Score.CheckDiscriminator()
	Score.CheckFlags()
	Score.CheckSnowflake()

	Score.Total = Score.Badges + Score.Discriminator + Score.Snowflake + Score.Avatar + Score.Banner
	return Score
}
