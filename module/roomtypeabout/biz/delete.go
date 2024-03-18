package roomtypeaboutbiz

import (
	"context"
)

type DeleteRoomTypeAboutStore interface {
	Delete(ctx context.Context, id int) error
}

type deleteRoomTypeAboutBiz struct {
	store DeleteRoomTypeAboutStore
}

func NewDeleteRoomTypeAboutBit(store DeleteRoomTypeAboutStore) *deleteRoomTypeAboutBiz {
	return &deleteRoomTypeAboutBiz{store: store}
}

func (biz *deleteRoomTypeAboutBiz) DeleteRoomTypeAbout(ctx context.Context,
	id int) error {

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
