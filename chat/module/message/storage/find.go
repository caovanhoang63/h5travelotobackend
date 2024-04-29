package chatmessagestorage

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) FindMessageWithCondition(ctx context.Context,
	condition map[string]interface{},
) (*chatmessagemodel.Message, error) {
	var result chatmessagemodel.Message
	coll := s.db.Collection(chatmessagemodel.Message{}.CollectionName())
	filter, err := bson.Marshal(condition)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	one := coll.FindOne(ctx, filter)
	err = one.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.ErrInternal(err)
		}
		return nil, common.ErrDb(err)
	}
	err = one.Decode(&result)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &result, nil
}
