package bookingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	bookingmodel "h5travelotobackend/module/bookings/model"
	"log"
)

func (biz *createBookingBiz) CreateFrontDeskBooking(ctx context.Context, data *bookingmodel.FrontDeskBookingCreate) error {
	if err := biz.ValidateBooking(ctx, data.Booking); err != nil {
		return err
	}

	err := biz.bookingStore.CreateFrontDeskBooking(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	var dtoBooking common.DTOBooking

	dtoBooking.Id = data.Booking.Id

	err = biz.pb.Publish(ctx, common.TopicCreateBooking, pubsub.NewMessage(dtoBooking))
	if err != nil {
		log.Println(common.ErrCannotCreateEntity(bookingmodel.EntityName, err))
	}
	return nil
}
