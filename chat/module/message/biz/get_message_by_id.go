package chatmessagebiz

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
)

type FindMessageStore interface {
	FindMessageWithCondition(ctx context.Context,
		condition map[string]interface{},
	) (*chatmessagemodel.Message, error)
}

type findMessageBiz struct {
	store FindMessageStore
}

func NewFindMessageBiz(store FindMessageStore) *findMessageBiz {
	return &findMessageBiz{store: store}
}

func (biz *findMessageBiz) GetMessageById(ctx context.Context, id *primitive.ObjectID,
	requester common.Requester) (*chatmessagemodel.Message, error) {

	cond := map[string]interface{}{"_id": id}

	message, err := biz.store.FindMessageWithCondition(ctx, cond)

	if err != nil {
		if errors.Is(err, common.DocumentNotFound) {
			return nil, common.ErrEntityNotFound(chatmessagemodel.EntityName, err)
		}
		return nil, common.ErrInternal(err)
	}

	return message, nil
}
