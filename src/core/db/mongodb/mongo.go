package mongodb

import (
	"context"
	"github.com/utsav-vaghani/video-search/src/common/configs"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

// NewConnection returns singleton instance of mongo db.
func NewConnection() (*mongo.Database, error) {
	if db != nil {
		return db, nil
	}

	cfg := configs.GetConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDB on %s/%s\n", cfg.MongoURI, cfg.DBName)

	db = client.Database(cfg.DBName)

	return db, nil
}
