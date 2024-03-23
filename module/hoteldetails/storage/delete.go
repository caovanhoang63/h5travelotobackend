package hoteldetailstorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

func (s *mongoStore) DeleteAdditionalInfo(ctx context.Context, id int) error {
	update := bson.D{{"$set", bson.D{{"status", 0}}}}
	_, err := s.db.Collection(hoteldetailmodel.HotelDetail{}.CollectionName()).
		UpdateOne(ctx, bson.D{{"hotel_id", id}}, update)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
