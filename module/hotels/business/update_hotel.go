package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type HotelUpdateSqlStore interface {
	Update(ctx context.Context, id int, update *hotelmodel.HotelUpdate) error
}

type hotelUpdateBiz struct {
	sqlStore  HotelUpdateSqlStore
	findRepo  FindHotelRepo
	requester common.Requester
}

func NewHotelUpdateBiz(sqlStore HotelUpdateSqlStore, findRepo FindHotelRepo, requester common.Requester) *hotelUpdateBiz {
	return &hotelUpdateBiz{
		sqlStore:  sqlStore,
		findRepo:  findRepo,
		requester: requester,
	}
}

func (b *hotelUpdateBiz) UpdateBase(ctx context.Context, id int, update *hotelmodel.HotelUpdate) error {
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

	if err := update.Validate(); err != nil {
		return err
	}

	err = b.sqlStore.Update(ctx, id, update)
	if err != nil {
		return common.ErrCannotUpdateEntity(hotelmodel.EntityName, err)
	}
	return nil
}
