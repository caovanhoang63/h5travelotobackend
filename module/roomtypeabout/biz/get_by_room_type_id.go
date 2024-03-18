package roomtypeaboutbiz

import (
	"context"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

type FindRoomTypeAboutStore interface {
	FindWithCondition(ctx context.Context, condition map[string]interface{}) (*roomtypeaboutmodel.RoomTypeAbout, error)
}

type findRoomTypeAboutBiz struct {
	store FindRoomTypeAboutStore
}

func NewFindRoomTypeAboutBiz(store FindRoomTypeAboutStore) *findRoomTypeAboutBiz {
	return &findRoomTypeAboutBiz{store: store}
}

func (biz *findRoomTypeAboutBiz) GetByRoomTypeId(ctx context.Context, roomTypeId int) (*roomtypeaboutmodel.RoomTypeAbout, error) {
	data, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"room_type_id": roomTypeId})
	if err != nil {
		return nil, err
	}
	return data, nil
}
