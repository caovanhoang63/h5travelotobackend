package bookingsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) GetHotelTotalRoom(ctx context.Context, hotelId int) (int, error) {
	var result int
	if err := s.db.Table("hotels").Where("id = ? ", hotelId).Pluck("total_room", &result).Error; err != nil {
		return 0, common.ErrDb(err)
	}
	return result, nil
}

func (s *sqlStore) OccupancyStatistic(ctx context.Context, hotelId int, date *common.CivilDate,
) (*bookingmodel.OccupancyStatistic, error) {
	var result bookingmodel.OccupancyStatistic
	err := s.db.Raw("SELECT "+
		"SUM(CASE WHEN DATE(start_date) <= DATE(?) AND DATE(end_date) >= DATE(?) THEN bookings.room_quantity ELSE 0 END) AS day_0, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 1 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 1 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_1, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 2 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 2 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_2, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 3 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 3 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_3, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 4 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 4 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_4, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 5 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 5 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_5, "+
		"SUM(CASE WHEN DATE(start_date) <= DATE_SUB(DATE(?), INTERVAL 6 DAY) AND DATE(end_date) >= DATE_SUB(DATE(?), INTERVAL 6 DAY) THEN bookings.room_quantity ELSE 0 END)  AS day_6 "+
		"FROM bookings "+
		"WHERE hotel_id = ? AND state NOT IN ('expired', 'canceled', 'deleted')",
		date, date, date, date, date, date, date, date, date, date, date, date, date, date, hotelId).Find(&result).Error
	if err != nil {
		return nil, common.ErrDb(err)
	}
	return &result, nil
}

func (s *sqlStore) RoomStatus(ctx context.Context, hotelId int, date *common.CivilDate) (*bookingmodel.RoomStatus, error) {
	var result bookingmodel.RoomStatus
	err := s.db.Raw("SELECT +"+
		"   (SELECT total_room "+
		"     FROM hotels "+
		"     WHERE id = ? ) AS total, "+
		"    (SELECT SUM(bookings.room_quantity) "+
		"     FROM bookings "+
		"     WHERE hotel_id = ? "+
		"        AND DATE(start_date) <= DATE(?) "+
		"        AND DATE(end_date) >= DATE(?) "+
		"        AND state NOT IN ('expired', 'canceled', 'deleted') "+
		"           AND status = 1 ) "+
		"  AS booked, "+
		"    fixing_and_dirty_rooms.fixing, "+
		"    fixing_and_dirty_rooms.dirty "+
		"FROM "+
		"    (SELECT "+
		"        SUM(CASE WHEN `condition` = 'fixing' THEN 1 ELSE 0 END) AS fixing, "+
		"        SUM(CASE WHEN `condition` = 'dirty' THEN 1 ELSE 0 END) AS dirty "+
		"     FROM rooms where hotel_id = ?) AS fixing_and_dirty_rooms;"+
		"", hotelId, hotelId, date, date, hotelId).First(&result).Error
	if err != nil {
		return nil, common.ErrDb(err)
	}

	return &result, nil
}

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
