package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

type RoomTypeStore interface {
	Create(ctx context.Context, data *roomtypemodel.RoomTypeCreate) error
}

type createRoomTypeBiz struct {
	store RoomTypeStore
}

func NewRoomTypeBiz(store RoomTypeStore) *createRoomTypeBiz {
	return &createRoomTypeBiz{store: store}
}

func (biz *createRoomTypeBiz) CreateRoomType(ctx context.Context, data *roomtypemodel.RoomTypeCreate) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(roomtypemodel.EntityName, err)
	}
	return nil
}
