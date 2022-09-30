package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	db *mongo.Collection
)

func Connect() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("[ERROR]: Cannot load environment variables - ", err)
	}

	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		log.Fatal("[ERROR]: Must set MONGO_URI")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("[ERROR]: Mongo connect failed ", err)
		panic(err)
	}
	db = client.Database("moviestore").Collection("movies")
	log.Print("[INFO] - MongoDB connection established successfully")
}

func GetDatabase() *mongo.Collection {
	return db
}
