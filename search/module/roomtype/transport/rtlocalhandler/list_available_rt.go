package rtSearchlocalhdl

import (
	"golang.org/x/net/context"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/asyncjob"
	"h5travelotobackend/module/bookings/transport/bklocalhandler"
	rtsearchbiz "h5travelotobackend/search/module/roomtype/biz"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	rtsearchrepo "h5travelotobackend/search/module/roomtype/repo"
	rtsearchstorage "h5travelotobackend/search/module/roomtype/storage"
	"h5travelotobackend/search/utils"
	"log"
)

type listAvailableRoomTypeHandler struct {
	appCtx appContext.AppContext
}

func (l *listAvailableRoomTypeHandler) ListAvailableRt(ctx context.Context,
	filter *rtsearchmodel.Filter,
) ([]rtsearchmodel.RoomType, error) {
	filter.SetDefault()
	rtHandlerStore := rtsearchstorage.NewStore(l.appCtx.GetElasticSearchClient(), l.appCtx.GetRedisClient())
	bookingHandler := bklocalhandler.NewCountBookedRoomLocalHandler(l.appCtx)
	rtHandlerRepo := rtsearchrepo.NewListRoomTypeRepo(rtHandlerStore, bookingHandler)
	rtHandler := rtsearchbiz.NewListAvailableRtBiz(rtHandlerRepo)

	rts, err := rtHandler.ListAvailableRt(ctx, filter)
	if err != nil {
		return nil, err
	}

	if rts == nil {
		return nil, nil
	}

	for i := range rts {
		rts[i].Mask(false)
		log.Println("Room type: ", rts[i].FakeId)
	}
	filter.Mask(false)

	//Cache available room types
	if err = asyncjob.NewJob(func(ctx context.Context) error {
		key := utils.GenCacheKeyForQuery(filter.QueryTime, filter.HotelId)
		err := rtHandlerRepo.CacheRoomTypes(ctx, key, rts, filter)
		if err != nil {
			return err
		}

		log.Println("cache key: ", key)
		return nil
	}).Execute(ctx); err != nil {
		log.Println("Error while caching room types: ", err)
	}

	return rts, nil
}

func NewListAvailableRoomTypeHandler(appCtx appContext.AppContext) *listAvailableRoomTypeHandler {
	return &listAvailableRoomTypeHandler{
		appCtx: appCtx,
	}
}
