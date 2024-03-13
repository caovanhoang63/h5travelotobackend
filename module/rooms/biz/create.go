package roombiz

import (
	"context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

type CreateRoomStore interface {
	Create(ctx context.Context, create *roommodel.RoomCreate) error
}

type createRoomBiz struct {
	store CreateRoomStore
}

func NewCreateRoomBiz(store CreateRoomStore) *createRoomBiz {
	return &createRoomBiz{store: store}
}

func (biz *createRoomBiz) CreateRoom(ctx context.Context, data *roommodel.RoomCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(roommodel.EntityName, err)
	}

	return nil
}
