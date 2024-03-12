package hotelmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *mongoStore) Update(ctx context.Context, id int, update *hotelmodel.HotelAdditionalInfo) error {
	bsonUpdate := bson.D{{"$set", update}}
	filter := bson.D{{"hotel_id", id}}
	if _, err := s.db.Collection(update.CollectionName()).UpdateOne(ctx, filter, bsonUpdate); err != nil {
		return common.ErrDb(err)
	}

	return nil
}
