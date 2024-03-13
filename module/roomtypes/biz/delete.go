package roomtypebiz

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

type DeleteRoomTypeStore interface {
	Delete(ctx context.Context, id int) error
}

type deleteRoomTypeBiz struct {
	store DeleteRoomTypeStore
}

func NewDeleteRoomTypeBiz(store DeleteRoomTypeStore) *deleteRoomTypeBiz {
	return &deleteRoomTypeBiz{store: store}
}

func (biz *deleteRoomTypeBiz) DeleteRoomType(ctx context.Context, id int) error {

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(roomtypemodel.EntityName, err)
	}

	return nil

}
