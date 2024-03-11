package hotelmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *mongoStore) DeleteAdditionalInfo(ctx context.Context, id int) error {
	update := bson.D{{"$set", bson.D{{"status", 0}}}}
	_, err := s.db.Collection(hotelmodel.HotelAdditionalInfo{}.CollectionName()).
		UpdateOne(ctx, bson.D{{"hotel_id", id}}, update)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
