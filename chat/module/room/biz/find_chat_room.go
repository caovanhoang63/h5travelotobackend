package chatroombiz

import (
	"errors"
	"golang.org/x/net/context"
	chatroom "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

type FindChatRoomStore interface {
	FindChatRoom(ctx context.Context, userId int, hotelId int) (*chatroom.Room, error)
	CreateRoom(ctx context.Context, create *chatroom.RoomCreate) error
}

type findChatRoomBiz struct {
	store FindChatRoomStore
}

func NewFindChatRoomBiz(store FindChatRoomStore) *findChatRoomBiz {
	return &findChatRoomBiz{store: store}
}

func (biz *findChatRoomBiz) FindChatRoom(ctx context.Context,
	userId int, hotelId int) (*chatroom.Room, error) {
	room, err := biz.store.FindChatRoom(ctx, userId, hotelId)
	if err == nil {
		return room, nil
	}
	if errors.Is(err, common.DocumentNotFound) {
		roomCreate := &chatroom.RoomCreate{
			HotelId: hotelId,
			UserId:  userId,
		}
		if err = biz.store.CreateRoom(ctx, roomCreate); err != nil {
			return nil, common.ErrInternal(err)
		}
		room, err = biz.store.FindChatRoom(ctx, userId, hotelId)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		room.UserId = userId
		room.HotelId = hotelId
		return room, nil
	}

	return nil, common.ErrInternal(err)
}
