package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
	hotelfacilitysqlstore "h5travelotobackend/module/hotelfacilities/storage/sqlstore"
	"log"
)

func CreateHotelFacilityDetails(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "storage facilities of a hotel when it's created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var hotel common.DTOHotel
			err := json.Unmarshal(message.Data, &hotel)
			if err != nil {
				log.Println(err)
			}
			hotelFacilityDetails := make([]hotelfacilitymodel.HotelFacilityDetail, len(hotel.FacilitiesIds))

			for i := range hotel.FacilitiesIds {
				hotelFacilityDetails[i].HotelId = hotel.Id
				uid, err := common.FromBase58(hotel.FacilitiesIds[i])
				if err != nil {
					log.Println(err)
				}
				hotelFacilityDetails[i].FacilityId = int(uid.GetLocalID())
			}
			store := hotelfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
			return store.CreateHotelFacilityDetails(ctx, hotelFacilityDetails)
		},
	}
}
