package database

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	VerificationLevelShowNever                 = 0
	VerificationLevelShowAlways                = 1
	VerificationLevelShowSuspicious            = 2
	VerificationLevelShowSuspiciousAndBlockBot = 3
)

type Server struct {
	ID                       primitive.ObjectID `bson:"_id"`
	GuildID                  string             `bson:"guild_id"`
	VerificationLevel        int                `bson:"verification_level"`
	VerificationTime         int                `bson:"verification_time"`
	VerificationRoleID       string             `bson:"verification_role_id"`
	VerificationLogChannelID string             `bson:"verification_log_channel_id"`
}
