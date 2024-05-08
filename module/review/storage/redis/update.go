package reviewstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
)

func (s *store) IncTotalReview(ctx context.Context, hotelId int, point int) error {
	totalKey := reviewmodel.GetTotalKey(hotelId)
	avgKey := reviewmodel.GetAvgKey(hotelId)
	txPipeline := s.redis.TxPipeline()

	curTotalStr := txPipeline.Get(ctx, totalKey)
	total, err := curTotalStr.Int()
	if err != nil {
		return common.ErrInternal(err)
	}
	curAvgStr := txPipeline.Get(ctx, avgKey)
	curAvg, err := curAvgStr.Float32()
	if err != nil {
		return common.ErrInternal(err)
	}

	newAvg := (curAvg + float32(point)) / (float32(total) + float32(point))

	txPipeline.Incr(ctx, totalKey)
	txPipeline.Set(ctx, avgKey, newAvg, 0)
	_, err = txPipeline.Exec(ctx)
	if err != nil {
		return common.ErrInternal(err)
	}

	return nil
}

func (s *store) DecTotalReview(ctx context.Context, hotelId int, point int) error {
	totalKey := reviewmodel.GetTotalKey(hotelId)
	avgKey := reviewmodel.GetAvgKey(hotelId)
	txPipeline := s.redis.TxPipeline()

	curTotalStr := txPipeline.Get(ctx, totalKey)
	total, err := curTotalStr.Int()
	if err != nil {
		return common.ErrInternal(err)
	}
	curAvgStr := txPipeline.Get(ctx, avgKey)
	curAvg, err := curAvgStr.Float32()
	if err != nil {
		return common.ErrInternal(err)
	}

	newAvg := (curAvg - float32(point)) / (float32(total) - float32(point))

	txPipeline.Decr(ctx, totalKey)
	txPipeline.Set(ctx, avgKey, newAvg, 0)

	_, err = txPipeline.Exec(ctx)
	if err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
