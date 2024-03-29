package roomstorage

import (
	"context"
	"fmt"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

func (s *sqlStore) ListRoomWithCondition(
	ctx context.Context,
	filter *roommodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]roommodel.Room, error) {
	var data []roommodel.Room

	db := s.db.Table(roommodel.Room{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.HotelId > 0 {
			db = db.Where("hotel_id = ?", f.HotelId)
		}
		if f.RoomTypeId > 0 {
			fmt.Println("room type ", f.RoomTypeId)
			db = db.Where("room_type_id = ?", f.RoomTypeId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	// paging
	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("id  < ? ", uid.GetLocalID())
	} else {
		db = db.Offset(paging.GetOffSet())
	}

	if err := db.Limit(paging.Limit).Find(&data).Order("id desc").Error; err != nil {
		return nil, err
	}

	if len(data) > 0 {
		last := data[len(data)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return data, nil
}

func (s *sqlStore) ListRoomsNotInIds(
	ctx context.Context,
	condition map[string]interface{},
	ids []int,
) ([]roommodel.Room, error) {
	var data []roommodel.Room

	db := s.db.Table(roommodel.Room{}.TableName()).Where(condition)
	if len(ids) > 0 {
		db = db.Where("id NOT IN ?", ids)
	}
	if err := db.Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return data, nil
}
