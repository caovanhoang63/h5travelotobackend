package reviewstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) GetTotalAndAvg(ctx context.Context, hotelId int) (int, float32, error) {
	txPipeline := s.redis.TxPipeline()

	rvStr := txPipeline.Get(ctx, reviewmodel.GetTotalReviewKey(hotelId))
	rtStr := txPipeline.Get(ctx, reviewmodel.GetTotalRatingKey(hotelId))
	_, err := txPipeline.Exec(ctx)
	if err != nil {
		return 0, 0, common.ErrDb(err)
	}

	rv, err := rvStr.Int()
	if err != nil {
		return 0, 0, common.ErrInternal(err)
	}
	rt, err := rtStr.Float32()
	if err != nil {
		return 0, 0, common.ErrInternal(err)
	}

	avg := rt / float32(rv)
	return rv, avg, nil
}
