package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) Create(ctx context.Context, review *reviewmodel.Review) error {

	coll := s.db.Collection(review.CollectionName())
	one, err := coll.InsertOne(ctx, review)
	if one != nil {
		id := (one.InsertedID).(primitive.ObjectID)
		review.ID = &id
		fmt.Println("review")
	}
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
