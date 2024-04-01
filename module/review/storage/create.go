package reviewstorage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *mongoStore) Create(ctx context.Context, review *reviewmodel.Review) error {

	coll := s.db.Collection(review.CollectionName())
	one, err := coll.InsertOne(ctx, review)
	review.ID = (one.InsertedID).(primitive.ObjectID)
	fmt.Println("review")
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}