package bookingsqlstorage

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) ListBookingWithCondition(
	ctx context.Context,
	filter *bookingmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]bookingmodel.Booking, error) {
	var result []bookingmodel.Booking

	db := s.db.Table(bookingmodel.Booking{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if filter.UserId > 0 {
			db = db.Where("user_id = ?", f.UserId)
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
