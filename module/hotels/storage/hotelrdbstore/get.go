package hotelrdbstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

func (s *store) GetHotelCurrentlyViewed(ctx context.Context, userId int) ([]int, error) {
	var hotelIds []int
	key := common.GenUserRecentViewedKey(userId)

	err := s.rdb.LRange(ctx, key, 0, 20).ScanSlice(&hotelIds)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	return hotelIds, nil
}
