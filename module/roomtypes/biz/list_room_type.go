package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

type ListRoomTypeStore interface {
	ListRoomTypeWithCondition(
		ctx context.Context,
		filter *roomtypemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]roomtypemodel.RoomType, error)
}

type listRoomTypeBiz struct {
	store ListRoomTypeStore
}

func NewListRoomTypeBiz(store ListRoomTypeStore) *listRoomTypeBiz {
	return &listRoomTypeBiz{store: store}
}

func (biz *listRoomTypeBiz) ListRoomTypeWithCondition(
	ctx context.Context,
	filter *roomtypemodel.Filter,
	paging *common.Paging) ([]roomtypemodel.RoomType, error) {

	data, err := biz.store.ListRoomTypeWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(roomtypemodel.EntityName, err)
	}

	return data, nil
}
