package hotelsearchrdbstore

import (
	"golang.org/x/net/context"
	reviewmodel "h5travelotobackend/module/review/model"
	"math"
)

func (s *store) GetHotelRating(ctx context.Context, id int) (float64, int, error) {
	tx := s.rdb.TxPipeline()
	ratingCMD := tx.Get(ctx, reviewmodel.GetTotalRatingKey(id))
	reviewCMD := tx.Get(ctx, reviewmodel.GetTotalReviewKey(id))
	_, err := tx.Exec(ctx)
	if err != nil {
		return 0, 0, nil
	}
	rating, err := ratingCMD.Int()

	if err != nil {
		return 0, 0, nil
	}
	review, err := reviewCMD.Int()
	if err != nil {
		return 0, 0, nil
	}
	avg := float64(rating) / float64(review)

	return toFixed(avg, 1), review, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
