package roombiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
	roommodel "h5travelotobackend/module/rooms/model"
)

type DeleteRoomStore interface {
	Delete(ctx context.Context, id int) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*roommodel.Room, error)
}

type deleteRoomBiz struct {
	store DeleteRoomStore
}

func NewDeleteRoomBiz(store DeleteRoomStore) *deleteRoomBiz {
	return &deleteRoomBiz{store: store}
}

func (biz *deleteRoomBiz) DeleteRoom(ctx context.Context, id int) error {
	if oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrEntityNotFound(roommodel.EntityName, err)
	} else {
		if oldData.Status == 0 {
			return common.ErrEntityDeleted(roommodel.EntityName, nil)
		}
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	return nil
}
