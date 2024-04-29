package chatroomstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) CreateRoom(ctx context.Context, create *chatroommodel.RoomCreate) error {
	create.OnCreate()
	coll := s.db.Collection(chatroommodel.Room{}.CollectionName())
	_, err := coll.InsertOne(ctx, create)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
