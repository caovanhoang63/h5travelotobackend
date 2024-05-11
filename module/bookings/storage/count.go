package bookingsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) CountBookedRoom(ctx context.Context, rtId int,
	startDate *common.CivilDate, endDate *common.CivilDate) (*int, error) {
	var result int64
	db := s.db.Table(bookingmodel.Booking{}.TableName())

	if err := db.Count(&result).Where("room_type_id = ?", rtId).
		Where("bookings.end_date >= ? and  bookings.end_date <= ? ", startDate, endDate).
		Or("bookings.start_date >= ? and bookings.start_date <= ?", startDate, endDate).
		Error; err != nil {
		return nil, common.ErrDb(err)
	}

	var resultInt = int(result)
	return &resultInt, nil
}
