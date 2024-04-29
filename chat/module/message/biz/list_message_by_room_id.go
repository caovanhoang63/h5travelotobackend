package chatmessagebiz

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
)

type ListMessageStore interface {
	ListMessageWithCondition(ctx context.Context,
		filter *chatmessagemodel.Filter, paging *common.Paging) ([]chatmessagemodel.Message, error)
}

type listMessageBiz struct {
	store ListMessageStore
}

func NewListMessageBiz(store ListMessageStore) *listMessageBiz {
	return &listMessageBiz{store: store}
}

func (biz *listMessageBiz) ListMessageByRoomId(ctx context.Context,
	roomId *primitive.ObjectID, paging *common.Paging) ([]chatmessagemodel.Message, error) {

	filter := chatmessagemodel.Filter{RoomId: roomId}
	data, err := biz.store.ListMessageWithCondition(ctx, &filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return data, nil
}
