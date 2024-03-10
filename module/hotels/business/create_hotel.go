package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type CreateHotelRepo interface {
	Create(ctx context.Context, data *hotelmodel.HotelCreate) error
}

type createHotelBiz struct {
	repo CreateHotelRepo
}

func NewCreateHotelBiz(repo CreateHotelRepo) *createHotelBiz {
	return &createHotelBiz{repo: repo}
}

func (biz *createHotelBiz) CreateHotel(ctx context.Context, data *hotelmodel.HotelCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(hotelmodel.EntityName, err)
	}

	return nil
}
