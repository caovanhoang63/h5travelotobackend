package chatmessagebiz

import (
	"golang.org/x/net/context"
	chatmodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	"log"
)

type CreateNewMessageStore interface {
	CreateMessage(ctx context.Context, create *chatmodel.Message) error
}

type CreateNewMessageBiz struct {
	store CreateNewMessageStore
	ps    pubsub.Pubsub
}

func NewCreateNewMessageBiz(store CreateNewMessageStore,
	ps pubsub.Pubsub) *CreateNewMessageBiz {
	return &CreateNewMessageBiz{store: store, ps: ps}
}

func (biz *CreateNewMessageBiz) CreateMessage(ctx context.Context,
	create *chatmodel.Message,
) error {
	create.OnCreate()
	if err := biz.store.CreateMessage(ctx, create); err != nil {
		return common.ErrInternal(err)
	}
	message := pubsub.NewMessage(create)
	message.SetChannel(common.EventNewMessage)
	if err := biz.ps.Publish(ctx, common.EventNewMessage, message); err != nil {
		log.Println(common.ErrCannotPublishMessage(common.EventNewMessage, err))
	}
	return nil
}
