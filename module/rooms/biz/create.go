package roombiz

import (
	"context"
	"errors"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	roommodel "h5travelotobackend/module/rooms/model"
)

type CreateRoomStore interface {
	Create(ctx context.Context, create *roommodel.RoomCreate) error
}

type FindRoomTypeStore interface {
	FindDTODataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.DTORoomType, error)
}

type createRoomBiz struct {
	store     CreateRoomStore
	findStore FindRoomTypeStore
	pb        pubsub.Pubsub
}

func NewCreateRoomBiz(store CreateRoomStore, findStore FindRoomTypeStore, pb pubsub.Pubsub) *createRoomBiz {
	return &createRoomBiz{store: store, findStore: findStore, pb: pb}
}

func (biz *createRoomBiz) CreateRoom(ctx context.Context, data *roommodel.RoomCreate) error {
	roomType, err := biz.findStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": data.RoomTypeID})
	if err != nil {
		return common.ErrCannotCreateEntity(roommodel.EntityName, err)
	}

	if roomType.HotelId != data.HotelId || roomType.Id != data.RoomTypeID {
		return common.ErrCannotCreateEntity(roommodel.EntityName, errors.New("room type not exist"))
	}

	mess := pubsub.NewMessage(roomType)
	mess.SetChannel(common.TopicCreateNewRoom)
	err = biz.pb.Publish(ctx, common.TopicCreateNewRoom, mess)
	if err != nil {
		return common.NewCustomError(
			err,
			"cannot increase total room",
			"CANNOT_INCREASE_TOTAL_ROOM")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(roommodel.EntityName, err)
	}

	return nil
}
