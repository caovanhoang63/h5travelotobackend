package roomfacilitystorage

import (
	"context"
	"h5travelotobackend/common"
	roomfacilitiesmodel "h5travelotobackend/module/roomfacilities/model"
)

func (s *mongoStore) ListAllRoomFacilities(ctx context.Context) ([]roomfacilitiesmodel.RoomFacility, error) {
	coll := s.db.Collection(roomfacilitiesmodel.RoomFacility{}.CollectionName())
	find, err := coll.Find(ctx, nil)
	if err != nil {
		return nil, common.ErrDb(err)
	}
	var data []roomfacilitiesmodel.RoomFacility
	err = find.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
