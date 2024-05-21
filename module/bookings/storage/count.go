package bookingsqlstorage

import (
	"database/sql"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
	"log"
)

func (s *sqlStore) CountBookedRoom(ctx context.Context, rtId int,
	startDate *common.CivilDate, endDate *common.CivilDate) (*int, error) {
	var result sql.NullInt64
	db := s.db.Table(bookingmodel.Booking{}.TableName())

	row := db.Where("room_type_id = ? ", rtId).
		Where("(start_date between (?) and (?)) or (end_date between (?) and (?))",
			startDate, endDate, startDate, endDate).
		Not("state IN (?) OR"+
			" (state = 'pending' AND TIMESTAMPDIFF(HOUR, created_at, NOW()) >= 1 )",
			[]string{"canceled", "deleted", "expired"}).
		Select("sum(room_quantity)").Row()

	if row.Err() != nil {
		return nil, common.ErrDb(row.Err())
	}

	err := row.Scan(&result)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	if !result.Valid {
		return nil, nil
	}

	log.Println("result: ", result)
	var resultInt = int(result.Int64)
	return &resultInt, nil
}
