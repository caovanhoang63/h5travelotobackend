package roomtypeaboutbiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

type CreateRoomTypeAboutStore interface {
	Create(ctx context.Context, data *roomtypeaboutmodel.RoomTypeAbout) error
}

type FindRoomTypeStore interface {
	FindDTODataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.DTORoomType, error)
}

type createRoomTypeAboutBiz struct {
	store          CreateRoomTypeAboutStore
	findStore      FindRoomTypeStore
	findAboutStore FindRoomTypeAboutStore
}

func NewCreateRoomTypeAboutBit(store CreateRoomTypeAboutStore, findAboutStore FindRoomTypeAboutStore, findStore FindRoomTypeStore) *createRoomTypeAboutBiz {
	return &createRoomTypeAboutBiz{store: store, findAboutStore: findAboutStore, findStore: findStore}
}

func (biz *createRoomTypeAboutBiz) CreateRoomTypeAbout(ctx context.Context,
	data *roomtypeaboutmodel.RoomTypeAbout) error {

	roomType, err := biz.findStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": data.RoomTypeId})
	if err != nil {
		return roomtypeaboutmodel.ErrInvalidRoomType
	}

	if roomType.Status == 0 {
		return roomtypeaboutmodel.ErrInvalidRoomType
	}

	findResult, _ := biz.findAboutStore.FindWithCondition(ctx, map[string]interface{}{"room_type_id": data.RoomTypeId})
	if findResult != nil {
		return roomtypeaboutmodel.ErrRoomTypeAboutExisted
	}

	if err := data.Validate(); err != nil {
		return err
	}
	data.OnCreate()

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
