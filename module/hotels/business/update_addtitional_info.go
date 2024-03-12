package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type HotelUpdateMongoStore interface {
	Update(ctx context.Context, id int, update *hotelmodel.HotelAdditionalInfo) error
}

type hotelAdditionalDataUpdateBiz struct {
	mongoStore HotelUpdateMongoStore
	findRepo   FindHotelRepo
	requester  common.Requester
}

func NewHotelAdditionalDataUpdateBiz(mongoStore HotelUpdateMongoStore, findRepo FindHotelRepo, requester common.Requester) *hotelAdditionalDataUpdateBiz {
	return &hotelAdditionalDataUpdateBiz{
		mongoStore: mongoStore,
		findRepo:   findRepo,
		requester:  requester,
	}
}

func (b *hotelAdditionalDataUpdateBiz) UpdateAdditionalInfo(ctx context.Context, id int, update *hotelmodel.HotelAdditionalInfo) error {
	result, err := b.findRepo.FindBaseDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(hotelmodel.EntityName, err)
	}

	if result.OwnerID != b.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	err = b.mongoStore.Update(ctx, id, update)
	if err != nil {
		return common.ErrCannotUpdateEntity(hotelmodel.EntityName, err)
	}
	return nil
}
