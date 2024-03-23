package bookingtrackingstorage

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

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

	fmt.Println("result", len(result))

	if len(result) > 0 {
		cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", result[len(result)-1].CreatedAt.Format(timeLayout))))
		paging.NextCursor = cursorStr
	}

	return result, nil

}
