package database

import (
	"FinalProject/configs"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDb(c configs.ProgrammingConfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(c.DbMongoUrl)

	// Membuat klien MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Error("Error connecting to MongoDB Atlas: %v", err)
		return nil, err
	}

	// Memeriksa koneksi ke MongoDB Atlas
	err = client.Ping(context.Background(), nil)
	if err != nil {
		logrus.Error("Error pinging MongoDB Atlas: %v", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB Atlas!")

	// Memilih database
	database := client.Database(c.DbMongoName)

	return database, nil
}
