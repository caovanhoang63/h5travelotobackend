package chatroombiz

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

type HotelSeenMessageStore interface {
	HotelSeenMessages(ctx context.Context, roomId *primitive.ObjectID) error
}

type hotelSeenMessageBiz struct {
	store HotelSeenMessageStore
}

func NewHotelSeenMessageBiz(store HotelSeenMessageStore) *hotelSeenMessageBiz {
	return &hotelSeenMessageBiz{store: store}
}

func (biz *hotelSeenMessageBiz) HotelSeenMessages(ctx context.Context, roomId *primitive.ObjectID) error {
	if err := biz.store.HotelSeenMessages(ctx, roomId); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
