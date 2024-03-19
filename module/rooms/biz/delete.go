package roombiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
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
	pb    pubsub.Pubsub
}

func NewDeleteRoomBiz(store DeleteRoomStore, pb pubsub.Pubsub) *deleteRoomBiz {
	return &deleteRoomBiz{store: store, pb: pb}
}

func (biz *deleteRoomBiz) DeleteRoom(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(roommodel.EntityName, err)
	} else {
		if oldData.Status == 0 {
			return common.ErrEntityDeleted(roommodel.EntityName, nil)
		}
	}

	// Side effect : decrease room_type.total_room

	roomType := common.DTORoomType{Id: oldData.RoomTypeID}
	mess := pubsub.NewMessage(roomType)
	mess.SetChannel(common.TopicDeleteRoom)
	err = biz.pb.Publish(ctx, common.TopicDeleteRoom, mess)
	if err != nil {
		return common.NewCustomError(
			err,
			"cannot decrease total room",
			"CANNOT_DECREASE_TOTAL_ROOM")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	return nil
}
