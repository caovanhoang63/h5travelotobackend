package hotelrepo

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type DeleteHotelSqlStore interface {
	Delete(ctx context.Context, id int) error
}

type DeleteHotelMongoStore interface {
	DeleteAdditionalInfo(ctx context.Context, id int) error
}

type deleteHotelRepo struct {
	sqlStore   DeleteHotelSqlStore
	mongoStore DeleteHotelMongoStore
}

func NewDeleteHotelRepo(mongoStore DeleteHotelMongoStore, sqlStore DeleteHotelSqlStore) *deleteHotelRepo {
	return &deleteHotelRepo{
		sqlStore:   sqlStore,
		mongoStore: mongoStore,
	}
}

func (r *deleteHotelRepo) DeleteHotel(ctx context.Context, id int) error {

	if err := r.sqlStore.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	if err := r.mongoStore.DeleteAdditionalInfo(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	return nil
}
