package bookingsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) OverviewByDate(ctx context.Context, hotelId int, date *common.CivilDate,
) (*bookingmodel.BookingStatistic, error) {
	var result bookingmodel.BookingStatistic
	err := s.db.Raw(""+
		"SELECT "+
		"    total_checked_in,"+
		"    total_checked_out,"+
		"    total_in_hotel "+
		"FROM "+
		"    ("+
		"        SELECT "+
		"            SUM(CASE WHEN t.state = 'checked-in' THEN 1 ELSE 0 END) AS total_checked_in,"+
		"            SUM(CASE WHEN t.state = 'checked-out' THEN 1 ELSE 0 END) AS total_checked_out"+
		"        FROM "+
		"            booking_trackings t"+
		"        JOIN "+
		"          bookings b ON t.booking_id = b.id "+
		"        WHERE "+
		"            DATE(t.created_at) = DATE(?) AND b.hotel_id = ? "+
		"    ) AS trackings_summary, "+
		"    ( "+
		"        SELECT "+
		"            SUM(adults + children) AS total_in_hotel "+
		"        FROM "+
		"            bookings "+
		"        WHERE "+
		"            state = 'checked-in' AND hotel_id = ? "+
		"    ) AS bookings_summary;", date, hotelId, hotelId).First(&result).Error
	if err != nil {
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
