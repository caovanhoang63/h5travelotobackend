package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

func (biz *listBookingBiz) ListBookingByHotelId(
	ctx context.Context,
	filter *bookingmodel.Filter,
	paging *common.Paging,
) ([]bookingmodel.Booking, error) {

	result, err := biz.store.ListBookingWithCondition(ctx, filter, paging, "RoomType")
	if err != nil {
		return nil, common.ErrCannotListEntity(bookingmodel.EntityName, err)
	}

	for i := range result {
		ids, err := biz.bdStore.ListRoomIdsOfBooking(ctx, result[i].Id)
		if err == nil {
			result[i].Rooms = ids
		}
	}

	return result, nil
}
