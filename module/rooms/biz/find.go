package roombiz

import (
	"context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

type FindRoomStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*roommodel.Room, error)
}

type findRoomBiz struct {
	store FindRoomStore
}

func NewFindRoomBiz(store FindRoomStore) *findRoomBiz {
	return &findRoomBiz{store: store}
}

func (biz *findRoomBiz) GetRoomByID(ctx context.Context, id int) (*roommodel.Room, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrEntityNotFound(roommodel.EntityName, err)
	}

	return result, nil
}
