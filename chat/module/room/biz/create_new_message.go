package chatbiz

import (
	"golang.org/x/net/context"
	chatmodel "h5travelotobackend/chat/module/room/model/message"
)

type CreateNewMessageStore interface {
	CreateMessage(ctx context.Context, roomId string, create *chatmodel.Message) error
}

type CreateNewMessageBiz struct {
	store CreateNewMessageStore
}

func NewCreateNewMessageBiz(store CreateNewMessageStore) *CreateNewMessageBiz {
	return &CreateNewMessageBiz{store: store}
}

func (biz *CreateNewMessageBiz) CreateMessage(ctx context.Context, roomId string, create *chatmodel.Message) error {
	return biz.store.CreateMessage(ctx, roomId, create)
}
