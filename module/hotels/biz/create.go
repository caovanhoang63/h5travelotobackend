package hotelbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	hotelmodel "h5travelotobackend/module/hotels/model"
	"log"
)

type CreateHotelStore interface {
	Create(ctx context.Context, data *hotelmodel.HotelCreate) error
}

type createHotelBiz struct {
	store CreateHotelStore
	ps    pubsub.Pubsub
}

func NewCreateHotelBiz(store CreateHotelStore, ps pubsub.Pubsub) *createHotelBiz {
	return &createHotelBiz{
		store: store,
		ps:    ps,
	}
}

func (biz *createHotelBiz) CreateHotel(ctx context.Context, requester common.Requester, data *hotelmodel.HotelCreate) error {
	if requester.GetRole() != common.RoleOwner {
		return common.ErrNoPermission(errors.New("user is not owner"))
	}

	data.OwnerID = requester.GetUserId()

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(hotelmodel.EntityName, err)
	}
	mess := pubsub.NewMessage(&common.DTOHotel{Id: data.Id, OwnerId: data.OwnerID, FacilitiesIds: data.FacilityIds})

	if err := biz.ps.Publish(ctx, common.TopicCreateHotel, mess); err != nil {
		log.Println(common.ErrCannotPublishMessage(common.TopicCreateHotel, err))
	}

	return nil
}
