package hotelrdbstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

func (s *store) AddToUserRecentlyViewedHotel(ctx context.Context, userId, hotelId int) error {
	key := common.GenUserRecentViewedKey(userId)

	tx := s.rdb.TxPipeline()

	tx.LRem(ctx, key, 0, hotelId)
	tx.LPush(ctx, key, hotelId)
	tx.LTrim(ctx, key, 0, 20)

	_, err := tx.Exec(ctx)
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
