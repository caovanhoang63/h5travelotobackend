package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	reviewstorage "h5travelotobackend/module/review/storage/redis"
)

func UpdateHotelReviewWhenUserDeleteReview(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "update hotel review when user review hotel",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var review common.DTOReview

			err := json.Unmarshal(message.Data, &review)
			if err != nil {
				return err
			}

			store := reviewstorage.NewRedisStore(appCtx.GetRedisClient())

			return store.DecTotalReview(ctx, review.HotelId, review.Rating)
		},
	}
}
