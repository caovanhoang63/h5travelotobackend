package roombiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
	roommodel "h5travelotobackend/module/rooms/model"
)

type UpdateRoomStore interface {
	Update(ctx context.Context, id int, data *roommodel.RoomUpdate) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*roommodel.Room, error)
}

type updateRoomBiz struct {
	store UpdateRoomStore
}

func NewUpdateRoomBiz(store UpdateRoomStore) *updateRoomBiz {
	return &updateRoomBiz{store: store}
}

func (biz *updateRoomBiz) UpdateRoom(ctx context.Context, id int, data *roommodel.RoomUpdate) error {

	if oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrEntityNotFound(hotelmodel.EntityName, err)
	} else {
		if oldData.Status == 0 {
			return common.ErrEntityDeleted(hotelmodel.EntityName, nil)
		}
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(roommodel.EntityName, err)
	}

	return nil
}
