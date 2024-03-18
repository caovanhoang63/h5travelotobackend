package roomtypeaboutbiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

type UpdateRoomTypeAboutStore interface {
	Update(ctx context.Context, roomTypeId int, updateData *roomtypeaboutmodel.RoomTypeAboutUpdate) error
}

type updateRoomTypeAboutBiz struct {
	store UpdateRoomTypeAboutStore
}

func NewUpdateRoomTypeAboutBiz(store UpdateRoomTypeAboutStore) *updateRoomTypeAboutBiz {
	return &updateRoomTypeAboutBiz{store: store}
}

func (biz *updateRoomTypeAboutBiz) UpdateRoomTypeAbout(ctx context.Context, roomTypeId int, updateData *roomtypeaboutmodel.RoomTypeAboutUpdate) error {
	if err := biz.store.Update(ctx, roomTypeId, updateData); err != nil {
		return common.ErrCannotUpdateEntity(roomtypeaboutmodel.EntityName, err)
	}
	return nil
}
