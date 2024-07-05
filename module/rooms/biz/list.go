package roombiz

import (
	"context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

type ListRoomStore interface {
	ListRoomWithCondition(
		ctx context.Context,
		filter *roommodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]roommodel.Room, error)
	ListRoomsNotInIds(
		ctx context.Context,
		condition map[string]interface{},
		ids []int,
	) ([]roommodel.Room, error)
	ListRoomInIds(
		ctx context.Context,
		condition map[string]interface{},
		ids []int,
	) ([]roommodel.Room, error)
}

type listRoomBiz struct {
	store ListRoomStore
}

func NewListRoomBiz(store ListRoomStore) *listRoomBiz {
	return &listRoomBiz{store: store}
}

func (biz *listRoomBiz) ListRoomWithCondition(
	ctx context.Context,
	filter *roommodel.Filter,
	paging *common.Paging) ([]roommodel.Room, error) {

	data, err := biz.store.ListRoomWithCondition(ctx, filter, paging, "RoomType")

	if err != nil {
		return nil, common.ErrCannotListEntity(roommodel.EntityName, err)
	}

	return data, nil
}
