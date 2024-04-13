package dealsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

func (s *sqlStore) ListWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	filter *dealmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]dealmodel.Deal, error) {
	var result []dealmodel.Deal

	db := s.db.Table(dealmodel.Deal{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if filter.RoomTypeId > 0 {
			db = db.Where("room_type_id = ?", f.RoomTypeId)
		}
		if filter.HotelId > 0 {
			db = db.Where("hotel_id = ?", f.HotelId)
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

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}
	return result, nil
}
