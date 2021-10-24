package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Index struct {
	db *mongo.Database
}

func NewIndex(db *mongo.Database) *Index {
	return &Index{db: db}
}

func (s *Index) CreateIndexes(db *mongo.Database) {
	if err := s.UserIndex(); err != nil {
		log.Printf("%s", err.Error())
	}
}

func (s *Index) UserIndex() error {
	users := mongo.IndexModel{
		Keys: bson.M{
			"phone": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := s.db.Collection(usersCollection).Indexes().CreateOne(context.Background(), users)
	if err != nil {
		return err
	}
	return nil
}

func TitleSearchIndex(db *mongo.Database, collection string) error {
	adverts := mongo.IndexModel{
		Keys: bson.M{
			"title_search": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection(collection).Indexes().CreateOne(context.Background(), adverts)
	if err != nil {
		return err
	}
	return nil
}
