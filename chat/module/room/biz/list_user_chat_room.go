package chatroombiz

import (
	"golang.org/x/net/context"
	chatroommodel "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type listChatRoomByUserBiz struct {
	store  ListChatRoomStore
	hStore HotelStore
}

type HotelStore interface {
	ListByIds(ctx context.Context, ids []int) ([]hotelmodel.Hotel, error)
}

func NewListChatRoomByUserBiz(store ListChatRoomStore, hStore HotelStore) *listChatRoomByUserBiz {
	return &listChatRoomByUserBiz{store: store, hStore: hStore}
}

func (biz *listChatRoomByUserBiz) ListChatRoomByUser(ctx context.Context,
	requester common.Requester, paging *common.Paging) ([]chatroommodel.Room, error) {

	result, err := biz.store.ListChatRoom(ctx, &chatroommodel.Filter{UserId: requester.GetUserId()}, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var ids []int
	for i := range result {
		ids = append(ids, result[i].HotelId)
	}

	hotels, err := biz.hStore.ListByIds(ctx, ids)
	if err != nil {
		return result, nil
	}
	b := map[int]hotelmodel.Hotel{}
	for i := range hotels {
		b[hotels[i].Id] = hotels[i]
	}

	for i := range result {
		hotel := b[result[i].HotelId]
		result[i].Hotel = &hotel
	}
	return result, nil
}
