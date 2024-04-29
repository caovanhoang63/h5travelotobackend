package chatroombiz

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

type CustomerSeenMessageStore interface {
	UserSeenMessages(ctx context.Context, roomId *primitive.ObjectID) error
}

type customerSeenMessageBiz struct {
	store CustomerSeenMessageStore
}

func NewCustomerSeenMessageBiz(store CustomerSeenMessageStore) *customerSeenMessageBiz {
	return &customerSeenMessageBiz{store: store}
}

func (biz *customerSeenMessageBiz) CustomerSeenMessages(ctx context.Context, roomId *primitive.ObjectID) error {
	if err := biz.store.UserSeenMessages(ctx, roomId); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
