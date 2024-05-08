package mongo

import "go.mongodb.org/mongo-driver/mongo"

type store struct {
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *store {
	return &store{
		db: db,
	}
}
