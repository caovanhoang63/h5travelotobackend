package hotelmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *mongoStore) FindAdditionalInfo(ctx context.Context, id int) (*hotelmodel.HotelAdditionalInfo, error) {

	var data hotelmodel.HotelAdditionalInfo

	coll := s.db.Collection(hotelmodel.HotelAdditionalInfo{}.CollectionName())

	filter := bson.D{{"hotel_id", id}}

	result := coll.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	if err := result.Decode(&data); err != nil {
		return nil, common.ErrDb(err)
	}

	return &data, nil

}
