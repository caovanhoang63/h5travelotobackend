package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	workermodel "h5travelotobackend/module/worker/model"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
)

func CreateOwnerWorker(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Create owner worker when a hotel is created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var hotel common.DTOHotel
			if err := json.Unmarshal(message.Data, &hotel); err != nil {
				return err
			}
			db := appCtx.GetGormDbConnection()
			store := workersqlstorage.NewSqlStore(db)
			worker := workermodel.WorkerCreate{
				UserID:  hotel.OwnerId,
				HotelId: hotel.Id,
			}
			return store.Create(ctx, &worker)
		},
	}
}
