package chatmessagestorage

import (
	"golang.org/x/net/context"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) CreateMessage(ctx context.Context,
	create *chatmessagemodel.Message) error {

	coll := s.db.Collection(chatmessagemodel.Message{}.CollectionName())

	_, err := coll.InsertOne(ctx, create)

	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
