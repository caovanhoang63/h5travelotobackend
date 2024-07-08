package bookingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

type StatisticBookingStore interface {
	OverviewByDate(ctx context.Context, hotelId int, date *common.CivilDate,
	) (*bookingmodel.BookingStatistic, error)
	RoomStatus(ctx context.Context, hotelId int, date *common.CivilDate,
	) (*bookingmodel.RoomStatus, error)
	OccupancyStatistic(ctx context.Context, hotelId int, date *common.CivilDate,
	) (*bookingmodel.OccupancyStatistic, error)
	GetHotelTotalRoom(ctx context.Context, hotelId int) (int, error)
}
type statisticBookingBiz struct {
	store StatisticBookingStore
}

func NewStatisticBookingBiz(store StatisticBookingStore) *statisticBookingBiz {
	return &statisticBookingBiz{
		store: store,
	}
}

func (biz *statisticBookingBiz) OverviewByDate(ctx context.Context,
	hotelId int,
	date *common.CivilDate) (
	*bookingmodel.BookingStatistic, error) {

	data, err := biz.store.OverviewByDate(ctx, hotelId, date)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return data, nil

}

func (biz *statisticBookingBiz) OccupancyStatistic(ctx context.Context, hotelId int,
	date *common.CivilDate) (*bookingmodel.OccupancyStatistic, error) {
	data, err := biz.store.OccupancyStatistic(ctx, hotelId, date)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	total, err := biz.store.GetHotelTotalRoom(ctx, hotelId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	data.Day0 = data.Day0 / float64(total)
	data.Day1 = data.Day1 / float64(total)
	data.Day2 = data.Day2 / float64(total)
	data.Day3 = data.Day3 / float64(total)
	data.Day4 = data.Day4 / float64(total)
	data.Day5 = data.Day5 / float64(total)
	data.Day6 = data.Day6 / float64(total)

	return data, nil
}

func (biz *statisticBookingBiz) RoomStatus(ctx context.Context,
	hotelId int,
	date *common.CivilDate) (
	*bookingmodel.RoomStatus, error) {

	result, err := biz.store.RoomStatus(ctx, hotelId, date)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	result.Available = result.Total - result.Booked
	if result.Available < 0 {
		result.Available = 0
	}
	return result, nil
}
