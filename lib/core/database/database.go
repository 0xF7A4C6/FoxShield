package database

import (
	"bot/lib/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ServerCollection *mongo.Collection
	Ctx              = context.TODO()
)

func InitDatabase() {
	ClientOption := options.Client().ApplyURI(utils.DbConnectUrl)
	Client, Err := mongo.Connect(Ctx, ClientOption)

	if utils.HandleError(Err) {
		panic(Err)
	}

    Client.Connect(Ctx)
	ServerCollection = Client.Database("main").Collection("servers")
}

func AddGuild(GuildID string) {
	Count, Err := ServerCollection.CountDocuments(Ctx, bson.M{
		"guild_id": GuildID,
	})

	if utils.HandleError(Err) || Count != 0 {
		return
	}

    _, _ = ServerCollection.InsertOne(Ctx, Server{
        ID:                       primitive.NewObjectID(),
        GuildID:                  GuildID,
        VerificationLevel:        1,
        VerificationTime:         500,
        VerificationRoleID:       "",
        VerificationLogChannelID: "",
    })
}

func RemoveGuild(GuildID string) {
	ServerCollection.FindOneAndDelete(Ctx, bson.M{"guild_id": GuildID})
}

func GetGuild(GuildID string) Server {
	var Guild Server
	ServerCollection.FindOne(Ctx, bson.M{"guild_id": GuildID}).Decode(&Guild)
	return Guild
}

func UpdateGuild(GuildID string, ServerOption Server) {
	ServerCollection.FindOneAndUpdate(Ctx, bson.M{"guild_id": GuildID}, bson.M{
		"$set": bson.M{
			"verification_log_channel_id": ServerOption.VerificationLogChannelID,
			"verification_role_id":        ServerOption.VerificationRoleID,
			"verification_level":          ServerOption.VerificationLevel,
			"verification_time":           ServerOption.VerificationTime,
		},
	})
}
