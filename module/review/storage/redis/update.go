package reviewstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	reviewmodel "h5travelotobackend/module/review/model"
	"log"
)

func (s *store) IncTotalReview(ctx context.Context, hotelId int, point int) error {
	rvKey := reviewmodel.GetTotalReviewKey(hotelId)
	rtKey := reviewmodel.GetTotalRatingKey(hotelId)
	log.Println(rvKey)
	log.Println(rtKey)

	txPipeline := s.redis.TxPipeline()

	txPipeline.Incr(ctx, rvKey)
	txPipeline.IncrBy(ctx, rtKey, int64(point))

	_, err := txPipeline.Exec(ctx)
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}

func (s *store) DecTotalReview(ctx context.Context, hotelId int, point int) error {
	rvKey := reviewmodel.GetTotalReviewKey(hotelId)
	rtKey := reviewmodel.GetTotalRatingKey(hotelId)
	log.Println(rvKey)
	log.Println(rtKey)

	txPipeline := s.redis.TxPipeline()

	txPipeline.Incr(ctx, rvKey)
	txPipeline.IncrBy(ctx, rtKey, int64(point))

	_, err := txPipeline.Exec(ctx)
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
