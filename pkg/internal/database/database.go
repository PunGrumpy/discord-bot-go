package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type GuildConfig struct {
	GuildID     string `bson:"guildId"`
	DefaultRole string `bson:"defaultRole"`
}

const defaultRoleName = "UNVERIFY"

func NewDatabase(uri string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	collection := client.Database("your-database-name").Collection("your-collection-name")

	return &Database{
		Client:     client,
		Collection: collection,
	}, nil
}

func (db *Database) SetDefaultRole(guildID, roleID string) error {
	filter := bson.M{"guildId": guildID}
	update := bson.M{"$set": bson.M{"defaultRole": roleID}}

	_, err := db.Collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Printf("Error updating default role for guild %s: %v", guildID, err)
		return err
	}

	return nil
}

func (db *Database) GetDefaultRole(guildID string) (string, error) {
	filter := bson.M{"guildId": guildID}
	var config GuildConfig

	err := db.Collection.FindOne(context.TODO(), filter).Decode(&config)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If the document does not exist, return the default role name
			return defaultRoleName, nil
		}
		log.Printf("Error retrieving default role for guild %s: %v", guildID, err)
		return "", err
	}

	return config.DefaultRole, nil
}
