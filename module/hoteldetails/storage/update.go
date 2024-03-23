package hoteldetailstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

func (s *mongoStore) Update(ctx context.Context, id int, update *hoteldetailmodel.HotelDetail) error {
	bsonUpdate := bson.D{{"$set", update}}
	filter := bson.D{{"hotel_id", id}}
	if _, err := s.db.Collection(update.CollectionName()).UpdateOne(ctx, filter, bsonUpdate); err != nil {
		return common.ErrDb(err)
	}
	return nil
}
