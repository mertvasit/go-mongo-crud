package models

import (
	"context"
	"github.com/mertvasit/go-mongo-crud/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Movie struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Director string             `bson:"director"`
	Rating   float32            `bson:"rating"`
}

var db *mongo.Collection

func init() {
	config.Connect()
	db = config.GetDatabase()
}

func ReadMovies() ([]Movie, error) {
	var result []Movie
	filter := bson.M{}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := db.Find(ctx, filter)
	err = cur.All(ctx, &result)
	return result, err
}
func (m *Movie) CreateMovie() (*mongo.InsertOneResult, error) {
	m.Id = primitive.NewObjectID()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := db.InsertOne(ctx, m)
	return result, err
}
func ReadMovieById(Id primitive.ObjectID) (Movie, error) {
	var result Movie
	filter := bson.M{"_id": Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := db.FindOne(ctx, filter).Decode(&result)
	return result, err
}
func (m *Movie) UpdateMovie(Id primitive.ObjectID) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": Id}
	update := bson.D{{"$set", bson.M{
		"name":     m.Name,
		"director": m.Director,
		"rating":   m.Rating,
	}}}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := db.UpdateOne(ctx, filter, update)
	return result, err
}
func DeleteMovie(Id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": Id}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	mov, err := db.DeleteOne(ctx, filter)
	return mov, err
}
