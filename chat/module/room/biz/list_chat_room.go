package chatroombiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

type ListChatRoomStore interface {
	ListChatRoom(ctx context.Context, filter *chatroommodel.Filter, paging *common.Paging) ([]chatroommodel.Room, error)
}

type listChatRoomBiz struct {
	store ListChatRoomStore
}

func NewListChatRoomBiz(store ListChatRoomStore) *listChatRoomBiz {
	return &listChatRoomBiz{store: store}
}

func (biz *listChatRoomBiz) ListChatRoomByHotelId(ctx context.Context,
	hotelId int, paging *common.Paging) ([]chatroommodel.Room, error) {
	result, err := biz.store.ListChatRoom(ctx, &chatroommodel.Filter{HotelId: hotelId}, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
