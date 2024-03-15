package roombiz

import (
	"context"
	"errors"
	"h5travelotobackend/common"
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
	) (*common.RoomTypeDTO, error)
}

type createRoomBiz struct {
	store     CreateRoomStore
	findStore FindRoomTypeStore
}

func NewCreateRoomBiz(store CreateRoomStore, findStore FindRoomTypeStore) *createRoomBiz {
	return &createRoomBiz{store: store, findStore: findStore}
}

func (biz *createRoomBiz) CreateRoom(ctx context.Context, data *roommodel.RoomCreate) error {
	roomType, err := biz.findStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": data.RoomTypeID})
	if err != nil {
		return common.ErrCannotCreateEntity(roommodel.EntityName, err)
	}

	if roomType.HotelId != data.HotelId || roomType.Id != data.RoomTypeID {
		return common.ErrCannotCreateEntity(roommodel.EntityName, errors.New("room type not exist"))
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(roommodel.EntityName, err)
	}

	return nil
}
