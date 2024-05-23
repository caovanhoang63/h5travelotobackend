package bookingbiz

import (
	"golang.org/x/net/context"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

type checkDealBiz struct {
	dealStore DealStore
	rtStore   FindRoomTypeStore
}

func NewCheckDealBiz(dealStore DealStore, rtStore FindRoomTypeStore) *checkDealBiz {
	return &checkDealBiz{dealStore: dealStore, rtStore: rtStore}
}

func (biz *checkDealBiz) CheckDeal(ctx context.Context, bk *bookingmodel.BookingCreate, dealId int) error {
	deal, err := biz.dealStore.FindWithCondition(ctx, map[string]interface{}{"id": dealId})

	if err != nil {
		return bookingmodel.ErrInvalidDeal
	}

	roomType, err := biz.rtStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": bk.RoomTypeId})
	if err != nil {
		return bookingmodel.ErrInvalidRoomType
	}

	bk.TotalAmount = roomType.Price * float64(bk.RoomQuantity) * float64(bk.StartDate.DateDiff(*bk.EndDate))

	booking := bookingmodel.Booking{
		HotelId:     bk.HotelId,
		RoomTypeId:  bk.RoomTypeId,
		StartDate:   bk.StartDate,
		EndDate:     bk.EndDate,
		TotalAmount: bk.TotalAmount,
	}

	if err = ValidateDeal(&booking, deal); err != nil {
		return bookingmodel.ErrDealNotAvailable
	}

	err = CalculateDiscountAmount(deal, &booking)
	if err != nil {
		return bookingmodel.ErrDealNotAvailable
	}

	final := booking.TotalAmount - booking.DiscountAmount
	if final < 0 {
		final = 0
	}
	bk.DiscountAmount = booking.DiscountAmount
	bk.FinalAmount = final
	return nil
}
