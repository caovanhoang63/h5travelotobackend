package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func UpdateHotelReviewWhenUserReview(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "update hotel review when user review hotel",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var review common.DTOReview

			err := json.Unmarshal(message.Data, &review)
			if err != nil {
				return err
			}

			store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())

			return store.UpdateReviewWhenUserReview(ctx, &review)
		},
	}
}
