package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type DeleteHotelRepo interface {
	DeleteHotel(ctx context.Context, id int) error
}

type FindHotelBaseDataRepo interface {
	FindBaseDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*hotelmodel.Hotel, error)
}

type deleteHotelBiz struct {
	deleteRepo DeleteHotelRepo
	findRepo   FindHotelBaseDataRepo
	requester  common.Requester
}

func NewDeleteHotelBiz(deleteRepo DeleteHotelRepo, findRepo FindHotelBaseDataRepo, requester common.Requester) *deleteHotelBiz {
	return &deleteHotelBiz{deleteRepo: deleteRepo, findRepo: findRepo, requester: requester}
}

func (biz *deleteHotelBiz) DeleteHotel(ctx context.Context, id int) error {
	result, err := biz.findRepo.FindBaseDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(hotelmodel.EntityName, err)
	}

	if result.OwnerID != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.deleteRepo.DeleteHotel(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	return nil
}
