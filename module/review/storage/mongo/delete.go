package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) Delete(ctx context.Context, Id primitive.ObjectID) error {
	coll := s.db.Collection(reviewmodel.Review{}.CollectionName())

	filter := map[string]interface{}{"_id": Id}

	_, err := coll.UpdateOne(ctx, filter, map[string]interface{}{"$set": map[string]interface{}{"status": 0}})
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
