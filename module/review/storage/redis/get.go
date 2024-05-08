package reviewstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) GetTotalAndAvg(ctx context.Context, hotelId int) (int, float32, error) {
	txPipeline := s.redis.TxPipeline()

	totalStr := txPipeline.Get(ctx, reviewmodel.GetTotalKey(hotelId))
	total, err := totalStr.Int()
	if err != nil {
		return 0, 0, common.ErrInternal(err)
	}

	avgStr := txPipeline.Get(ctx, reviewmodel.GetAvgKey(hotelId))
	avg, err := avgStr.Float32()
	if err != nil {
		return 0, 0, common.ErrInternal(err)
	}

	return total, avg, nil
}
