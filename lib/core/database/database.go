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
	clientOptions := options.Client().ApplyURI("mongodb+srv://vichy:lmao@cluster0.hg7yr7j.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(Ctx, clientOptions)

	if utils.HandleError(err) {
		panic(err)
	}

	client.Connect(Ctx)
	ServerCollection = client.Database("main").Collection("servers")
}

func AddGuild(GuildID string) {
	Count, Err := ServerCollection.CountDocuments(Ctx, bson.M{
		"guild_id": GuildID,
	})

	if utils.HandleError(Err) || Count != 0 {
		return
	}

	ServerCollection.InsertOne(Ctx, Server{
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
