package chatstorage

import "go.mongodb.org/mongo-driver/mongo"

type mongoStore struct {
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *mongoStore {
	return &mongoStore{
		db: db,
	}
}
