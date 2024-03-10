package hotelmongostorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *mongoStore) Create(ctx context.Context, data *hotelmodel.HotelAdditionalInfo) error {
	coll := s.db.Collection(hotelmodel.HotelAdditionalInfo{}.CollectionName())
	if _, err := coll.InsertOne(ctx, data); err != nil {
		return common.ErrDb(err)
	}
	return nil
}
