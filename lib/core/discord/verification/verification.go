package verification

import (
	"bot/lib/core/database"
	"bot/lib/core/discord/algorithm"
	"bot/lib/utils"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	TaskList = map[string]Task{}
)

func XoR(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
			output += string(input[i] ^ key[i % len(key)])
	}

	return hex.EncodeToString([]byte(output))
}

func (T *Task) PoW() (string, string) {
	str := utils.RandomStr(60) + "." + fmt.Sprintf("%d", time.Now().Unix())
	out_str := ""

	out_str = XoR(str, T.TaskId)
	out_str = base64.RawURLEncoding.EncodeToString([]byte(out_str))

	return str, out_str
}

func (T *Task) CreateTaskChallenge() {
	In, Out := T.PoW()

	T.TaskChallenge = Challenge{
		InToken: In,
		OutResponse: Out,
	}
}

func CreateVerificationTask(UserID, GuildID string, User *discordgo.User) (Task, bool, error) {
	T := Task{
		MemberID:  UserID,
		GuildID:   GuildID,
		Passed:    false,
		Score:     algorithm.GetScore(User),
		TaskId:    utils.RandomStr(10),
		StartTime: time.Now(),
	}
	T.CreateTaskChallenge()

	TaskList[T.TaskId] = T

	Profil := fmt.Sprintf("[#%s] %s | Score: %d | Avatar: %d, Banner: %d, Discriminator: %d, Badges: %d, Snowflake: %d", T.TaskId, User, T.Score.Total, T.Score.Avatar, T.Score.Banner, T.Score.Discriminator, T.Score.Badges, T.Score.Snowflake)

	G := database.GetGuild(T.GuildID)

	if T.Score.Total < 15 {
		T.Behavior = BehaviorSpam
	} else if T.Score.Total > 80 {
		T.Behavior = BehaviorSafe
	} else if T.Score.Total < 40 {
		T.Behavior = BehaviorSuspicious
	} else {
		T.Behavior = BehaviorNormal
	}

	if G.VerificationLevel == database.VerificationLevelShowAlways {
		return T, false, nil
	}

	if G.VerificationLevel == database.VerificationLevelShowSuspiciousAndBlockBot {
		if T.Behavior == BehaviorSpam {
			utils.Debug("+", fmt.Sprintf("Blocked | %s", Profil))
			return T, false, fmt.Errorf("Your account has been flagged as suspicious, so it cannot pass verification")
		}

		utils.Debug("+", fmt.Sprintf("Captcha | %s", Profil))
		return T, false, nil
	}

	if G.VerificationLevel == database.VerificationLevelShowSuspicious {
		if T.Behavior == BehaviorSafe {
			utils.Debug("+", fmt.Sprintf("Passed | %s", Profil))
			return T, true, nil
		}

		return T, false, nil
	}

	if G.VerificationLevel == database.VerificationLevelShowNever {
		utils.Debug("+", fmt.Sprintf("Passed | %s", Profil))
		return T, true, nil
	}

	return T, false, nil
}

func (T Task) EndTask(S *discordgo.Session) {
	T.Passed = true

	G := database.GetGuild(T.GuildID)
	Err := S.GuildMemberRoleAdd(T.GuildID, T.MemberID, G.VerificationRoleID)

	L := ""
	EmbedTitle := ""
	Behavior := ""

	switch G.VerificationLevel {
	case database.VerificationLevelShowNever:
		L = "VerificationLevelShowNever"
	case database.VerificationLevelShowAlways:
		L = "VerificationLevelShowAlways"
	case database.VerificationLevelShowSuspicious:
		L = "VerificationLevelShowSuspicious"
	case database.VerificationLevelShowSuspiciousAndBlockBot:
		L = "VerificationLevelShowSuspiciousAndBlockBot"
	}

	switch T.Behavior {
	case BehaviorNormal:
		Behavior = "Normal"
	case BehaviorSafe:
		Behavior = "Safe"
	case BehaviorSpam:
		Behavior = "Spam"
	case BehaviorSuspicious:
		Behavior = "Suspicious"
	}

	if utils.HandleError(Err) {
		EmbedTitle = fmt.Sprintf("%s Verification failed.", utils.Emoji_red)
	} else {
		EmbedTitle = fmt.Sprintf("%s Verification passed.", utils.Emoji_green)
	}

	S.ChannelMessageSendEmbed(G.VerificationLogChannelID, &discordgo.MessageEmbed{
		Title: EmbedTitle,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  fmt.Sprintf("%s Username", utils.Emoji_aka),
				Value: fmt.Sprintf("`%s#%s`", T.Score.User.Username, T.Score.User.Discriminator),
			},
			{
				Name:  fmt.Sprintf("%s ID", utils.Emoji_profil),
				Value: fmt.Sprintf("`%s`", T.Score.User.ID),
			},
			{
				Name:  fmt.Sprintf("%s Score", utils.Emoji_fingerprint),
				Value: fmt.Sprintf("`%d`", T.Score.Total),
			},
			{
				Name:  fmt.Sprintf("%s  Task ID", utils.Emoji_todo),
				Value: fmt.Sprintf("`%s`", T.TaskId),
			},
			{
				Name:  fmt.Sprintf("%s Task Time", utils.Emoji_timeout),
				Value: fmt.Sprintf("`%fs`", time.Now().Sub(T.StartTime).Seconds()),
			},
			{
				Name:  fmt.Sprintf("%s Task Level", utils.Emoji_grow),
				Value: fmt.Sprintf("`%d - %s`", G.VerificationLevel, L),
			},
			{
				Name:  fmt.Sprintf("%s Behavior", utils.Emoji_warn),
				Value: fmt.Sprintf("`%d - %s`", T.Behavior, Behavior),
			},
		},
	})
}
