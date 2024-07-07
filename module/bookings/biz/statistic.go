package bookingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

type StatisticBookingStore interface {
	OverviewByDate(ctx context.Context, hotelId int, date *common.CivilDate,
	) (*bookingmodel.BookingStatistic, error)
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
