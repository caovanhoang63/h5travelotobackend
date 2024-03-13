package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

type FindRoomTypeBiz interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*roomtypemodel.RoomType, error)
}

type findRoomTypeBiz struct {
	store FindRoomTypeBiz
}

func NewFindRoomTypeBiz(store FindRoomTypeBiz) *findRoomTypeBiz {
	return &findRoomTypeBiz{store: store}
}

func (biz *findRoomTypeBiz) GetRoomTypeByID(ctx context.Context, id int) (*roomtypemodel.RoomType, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrEntityNotFound(roomtypemodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(roomtypemodel.EntityName, nil)
	}
	return result, nil

}
