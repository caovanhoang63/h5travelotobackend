package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
	"log"
)

type RoomTypeStore interface {
	Create(ctx context.Context, data *roomtypemodel.RoomTypeCreate) error
}

type createRoomTypeBiz struct {
	store RoomTypeStore
	ps    pubsub.Pubsub
}

func NewRoomTypeBiz(store RoomTypeStore, ps pubsub.Pubsub) *createRoomTypeBiz {
	return &createRoomTypeBiz{store: store, ps: ps}
}

func (biz *createRoomTypeBiz) CreateRoomType(ctx context.Context, data *roomtypemodel.RoomTypeCreate) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(roomtypemodel.EntityName, err)
	}

	dto := &common.DTORoomType{
		Id:          data.Id,
		Name:        data.Name,
		HotelId:     data.HotelId,
		Status:      data.Status,
		Price:       data.Price,
		FacilityIds: data.FacilityIds,
	}
	mess := pubsub.NewMessage(&dto)
	if err := biz.ps.Publish(ctx, common.TopicCreateRoomType, mess); err != nil {
		log.Println(common.ErrInvalidRequest(common.ErrCannotPublishMessage(common.TopicCreateRoomType, err)))
	}

	return nil
}
