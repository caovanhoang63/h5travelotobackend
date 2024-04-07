package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	hotelmodel "h5travelotobackend/module/hotels/model"
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

			hotel, err := store.FindDataWithCondition(ctx, map[string]interface{}{"id": review.HotelId})
			if err != nil {
				return err
			}

			rating := float32((hotel.Rating*float32(hotel.TotalRating) + float32(review.Rating)) / float32(hotel.TotalRating+1))
			totalRating := hotel.TotalRating + 1
			update := hotelmodel.HotelUpdate{
				Rating:      rating,
				TotalRating: totalRating,
			}

			return store.Update(ctx, review.HotelId, &update)
		},
	}
}
