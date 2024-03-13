package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

type RoomTypeUpdateStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*roomtypemodel.RoomType, error)
	Update(ctx context.Context, id int, update *roomtypemodel.RoomTypeUpdate) error
}

type roomTypeUpdateBiz struct {
	store RoomTypeUpdateStore
}

func NewRoomTypeUpdateBiz(store RoomTypeUpdateStore) *roomTypeUpdateBiz {
	return &roomTypeUpdateBiz{store: store}
}

func (biz *roomTypeUpdateBiz) Update(ctx context.Context, id int, data *roomtypemodel.RoomTypeUpdate) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(roomtypemodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(roomtypemodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(roomtypemodel.EntityName, err)
	}
	return nil
}
