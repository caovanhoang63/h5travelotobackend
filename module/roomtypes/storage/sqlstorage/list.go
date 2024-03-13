package roomtypesqlstorage

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

func (s *sqlStore) ListRoomTypeWithCondition(
	ctx context.Context,
	filter *roomtypemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]roomtypemodel.RoomType, error) {
	var data []roomtypemodel.RoomType

	db := s.db.Table(roomtypemodel.RoomType{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.HotelId > 0 {
			db = db.Where("hotel_id = ?", f.HotelId)
		}
		if f.FreeCancel {
			db = db.Where("free_cancel = ?", f.FreeCancel)
		}
		if f.BreakFast {
			db = db.Where("breakfast = ?", f.BreakFast)
		}
		if f.Bed != nil {
			if f.Bed.Single > 0 {
				db = db.Where("JSON_EXTRACT(bed, '$.single')  = ? ", f.Bed.Single)
			}
			if f.Bed.Double > 0 {
				db = db.Where("JSON_EXTRACT(bed, '$.double')  = ? ", f.Bed.Double)
			}
		}
		db.Where("price >= ? and price <= ?", f.MinPrice, f.MaxPrice)

		if f.StartDate != nil && f.EndDate != nil {
			//TODO: Xử lý tìm phòng trống trong thời gian yêu cầu
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
