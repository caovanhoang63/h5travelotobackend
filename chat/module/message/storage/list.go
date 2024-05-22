package chatmessagestorage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	chatmessagemodel "h5travelotobackend/chat/module/message/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) ListMessageWithCondition(ctx context.Context,
	filter *chatmessagemodel.Filter, paging *common.Paging) ([]chatmessagemodel.Message, error) {
	var result []chatmessagemodel.Message
	filterB, err := filter.ToBsonD()
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	filterB = append(filterB, bson.E{Key: "status", Value: common.StatusActive})

	option := &options.FindOptions{}

	if v := paging.FakeCursor; v == "" {
		option.SetSkip(int64(paging.GetOffSet()))
	} else {
		cursor := primitive.ObjectID{}
		err := cursor.UnmarshalText([]byte(v))
		if err != nil {
			filterB = append(filterB, bson.E{Key: "_id", Value: bson.M{"$lt": cursor}})
		} else {
			option.SetSkip(int64(paging.GetOffSet()))
		}
	}

	option.SetLimit(int64(paging.Limit)).SetSort(bson.D{{Key: "_id", Value: 1}})

	coll := s.db.Collection(chatmessagemodel.Message{}.CollectionName())
	if count, err := coll.CountDocuments(ctx, filterB); err != nil {
		return nil, common.ErrDb(err)
	} else {
		paging.Total = count
	}

	fmt.Println("filterB: ", filterB)
	fmt.Println("option: ", option)
	cur, err := coll.Find(ctx, filterB, option)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	for cur.Next(ctx) {
		var message chatmessagemodel.Message
		if err := cur.Decode(&message); err != nil {
			return nil, common.ErrDb(err)
		}
		result = append(result, message)
	}
	fmt.Println("count: ", cur.Current)

	if err := cur.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	if err := cur.Close(ctx); err != nil {
		return nil, common.ErrDb(err)
	}

	if len(result) > 0 {
		paging.NextCursor = result[len(result)-1].ID.Hex()
	}

	return result, nil
}
