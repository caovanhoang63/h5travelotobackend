package bookingdetailstorage

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) ListRoomsOfBooking(ctx context.Context, bookingId int) ([]bookingdetailmodel.Room, error) {
	var rooms []bookingdetailmodel.Room
	db := s.db
	if err := db.WithContext(ctx).
		Joins("right join booking_details on booking_details.room_id = rooms.id").
		Where("booking_id = ?", bookingId).
		Find(&rooms).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return rooms, nil
}

func (s *sqlStore) ListBookingRoomsWithCondition(
	conditions map[string]interface{},
	filter *bookingdetailmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]bookingdetailmodel.BookingDetail, error) {

	var result []bookingdetailmodel.BookingDetail

	db := s.db

	db = db.Table(bookingdetailmodel.BookingDetail{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.BookingId > 0 {
			db.Where("booking_id = ?", v.BookingId)
		}

		if v.RoomId > 0 {
			db.Where("room_id = ?", v.RoomId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDb(err)
		}
		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	if len(result) > 0 {
		cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", result[len(result)-1].CreatedAt.Format(timeLayout))))
		paging.NextCursor = cursorStr
	}

	return result, nil

}

func (s *sqlStore) CountRoomOfBooking(ctx context.Context, bookingId int) (int, error) {
	var count int64
	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName()).Where("booking_id = ?", bookingId)
	if err := db.Count(&count).Error; err != nil {
		return 0, common.ErrDb(err)
	}
	return int(count), nil
}

func (s *sqlStore) ListRoomIdsOfBooking(ctx context.Context, bookingId int) ([]int, error) {
	var ids []int

	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName()).Where("booking_id = ?", bookingId)
	if err := db.Order("room_id desc").Pluck("room_id", &ids).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return ids, nil
}

func (s *sqlStore) ListRoomOfBooking(ctx context.Context, bookingId int) ([]int, error) {
	var ids []int

	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName()).Where("booking_id = ?", bookingId)
	if err := db.Order("room_id desc").Pluck("room_id", &ids).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return ids, nil
}

func (s *sqlStore) ListRoomIdsBooked(ctx context.Context,
	startDate *time.Time,
	endDate *time.Time,
	condition map[string]interface{}) ([]int, error) {
	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName())
	var ids []int
	db = db.InnerJoins("JOIN bookings ON booking_details.booking_id = bookings.id")
	db = db.Where(condition)
	db = db.Where("bookings.status = ?", common.StatusActive)
	db = db.Where("bookings.end_date >= ? and  bookings.end_date <= ? ", startDate, endDate).
		Or("bookings.start_date >= ? and bookings.start_date <= ?", startDate, endDate)
	if err := db.Pluck("room_id", &ids).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return ids, nil
}
