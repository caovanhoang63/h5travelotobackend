package bookingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
	dealmodel "h5travelotobackend/module/deals/model"
)

type UpdateDeal interface {
	Update(ctx context.Context, id int, data *bookingmodel.BookingUpdate) error
	FindWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*bookingmodel.Booking, error)
}

type DealStore interface {
	FindWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*dealmodel.Deal, error)
}

type addDealBiz struct {
	bkStore   UpdateDeal
	dealStore DealStore
}

func NewAddDealBiz(bkStore UpdateDeal, dealStore DealStore) *addDealBiz {
	return &addDealBiz{bkStore: bkStore, dealStore: dealStore}
}

func (biz *addDealBiz) AddDeal(ctx context.Context, id int, dealId int) error {
	booking, err := biz.bkStore.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(bookingmodel.EntityName, err)
	}

	deal, err := biz.dealStore.FindWithCondition(ctx, map[string]interface{}{"id": dealId})
	if err != nil {
		return common.ErrEntityNotFound(dealmodel.EntityName, err)
	}

	if err = ValidateDeal(booking, deal); err != nil {
		return bookingmodel.ErrInvalidDeal
	}

	if err = CalculateDiscountAmount(deal, booking); err != nil {
		return bookingmodel.ErrInvalidDeal
	}

	final := booking.TotalAmount - booking.DiscountAmount
	update := bookingmodel.BookingUpdate{
		DealId:         &dealId,
		DiscountAmount: &booking.DiscountAmount,
		FinalAmount:    &final,
	}
	if err = biz.bkStore.Update(ctx, id, &update); err != nil {
		return common.ErrCannotUpdateEntity(bookingmodel.EntityName, err)
	}

	return nil

}

func ValidateDeal(b *bookingmodel.Booking, d *dealmodel.Deal) error {
	if d.Status == 0 {
		return bookingmodel.ErrDealNotAvailable
	}

	if d.HotelId != b.HotelId {
		return bookingmodel.ErrDealNotAvailable
	}

	if d.RoomTypeId != 0 && d.RoomTypeId != b.RoomTypeId {
		return bookingmodel.ErrDealNotAvailable
	}

	if d.StartDate.After(*b.StartDate) || d.ExpiryDate.Before(*b.EndDate) {
		return bookingmodel.ErrDealNotAvailable
	}
	if !d.IsUnlimited && d.TotalQuantity <= 0 {
		return bookingmodel.ErrDealNotAvailable
	}

	if b.TotalAmount < d.MinPrice {
		return bookingmodel.ErrDealNotAvailable
	}
	return nil
}

func CalculateDiscountAmount(d *dealmodel.Deal, b *bookingmodel.Booking) error {
	if d.DiscountType == "percent" {
		b.DiscountAmount = b.TotalAmount * d.DiscountPercent
	} else {
		b.DiscountAmount = d.DiscountAmount
	}
	return nil
}
