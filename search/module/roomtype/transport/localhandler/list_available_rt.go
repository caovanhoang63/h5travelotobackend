package rtSearchlocalHdl

import (
	"golang.org/x/net/context"
	"h5travelotobackend/component/appContext"
	rtsearchbiz "h5travelotobackend/search/module/roomtype/biz"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	rtsearchrepo "h5travelotobackend/search/module/roomtype/repo"
	rtsearchstorage "h5travelotobackend/search/module/roomtype/storage"
)

type listAvailableRoomTypeHandler struct {
	appCtx appContext.AppContext
}

func (l *listAvailableRoomTypeHandler) ListAvailableRt(ctx context.Context,
	filter *rtsearchmodel.Filter,
) ([]rtsearchmodel.RoomType, error) {
	rtHandlerStore := rtsearchstorage.NewStore(l.appCtx.GetElasticSearchClient())
	rtHandlerRepo := rtsearchrepo.NewListRoomTypeRepo(rtHandlerStore, nil)
	rtHandler := rtsearchbiz.NewListAvailableRtBiz(rtHandlerRepo)

	return rtHandler.ListAvailableRt(ctx, filter)
}

func NewListAvailableRoomTypeHandler(appCtx appContext.AppContext) *listAvailableRoomTypeHandler {
	return &listAvailableRoomTypeHandler{
		appCtx: appCtx,
	}
}
