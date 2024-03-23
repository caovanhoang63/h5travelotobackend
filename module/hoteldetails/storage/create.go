package hoteldetailstorage

import (
	"context"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

func (s *mongoStore) Create(ctx context.Context, data *hoteldetailmodel.HotelDetail) error {

	coll := s.db.Collection(data.CollectionName())

	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
