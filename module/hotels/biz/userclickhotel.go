package hotelbiz

import "golang.org/x/net/context"

type UserClickHotelStore interface {
	AddToUserRecentlyViewedHotel(ctx context.Context, userId, hotelId int) error
}

type userClickHotelBiz struct {
	store UserClickHotelStore
}

func NewUserClickHotelBiz(store UserClickHotelStore) *userClickHotelBiz {
	return &userClickHotelBiz{store: store}
}

func (biz *userClickHotelBiz) AddToUserRecentlyViewedHotel(ctx context.Context, userId, hotelId int) error {
	return biz.store.AddToUserRecentlyViewedHotel(ctx, userId, hotelId)
}
