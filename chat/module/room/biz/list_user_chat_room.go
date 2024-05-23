package chatroombiz

import (
	"golang.org/x/net/context"
	chatroommodel "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

type listChatRoomByUserBiz struct {
	store ListChatRoomStore
}

func NewListChatRoomByUserBiz(store ListChatRoomStore) *listChatRoomByUserBiz {
	return &listChatRoomByUserBiz{store: store}
}

func (biz *listChatRoomByUserBiz) ListChatRoomByUser(ctx context.Context,
	requester common.Requester, paging *common.Paging) ([]chatroommodel.Room, error) {
	result, err := biz.store.ListChatRoom(ctx, &chatroommodel.Filter{UserId: requester.GetUserId()}, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
