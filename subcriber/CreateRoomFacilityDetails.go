package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
	roomfacilitysqlstore "h5travelotobackend/module/roomfacilities/storage/sqlstore"
	"log"
)

func CreateRoomFacilityDetails(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "storage facilities of a Room when it's created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var Room common.DTORoomType
			err := json.Unmarshal(message.Data, &Room)
			if err != nil {
				log.Println(err)
			}
			RoomFacilityDetails := make([]roomfacilitymodel.RoomFacilityDetail, len(Room.FacilityIds))

			for i := range Room.FacilityIds {
				RoomFacilityDetails[i].RoomId = Room.Id
				RoomFacilityDetails[i].FacilityId = Room.FacilityIds[i]
			}
			store := roomfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
			return store.CreateHotelFacilityDetails(ctx, RoomFacilityDetails)
		},
	}
}
