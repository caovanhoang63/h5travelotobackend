package mongoStore

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

func (s *mongoStore) ListAllRoomFacilities(ctx context.Context) ([]hotelfacilitymodel.HotelFacility, error) {
	coll := s.db.Collection(hotelfacilitymodel.HotelFacility{}.CollectionName())
	find, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, common.ErrDb(err)
	}
	var data []hotelfacilitymodel.HotelFacility
	err = find.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
