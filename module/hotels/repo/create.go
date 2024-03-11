package hotelrepo

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type CreateHotelSqlStore interface {
	Create(ctx context.Context, data *hotelmodel.HotelCreate) error
}

type CreateHotelMongoStore interface {
	Create(ctx context.Context, data *hotelmodel.HotelAdditionalInfo) error
}

type createHotelRepo struct {
	sqlStore   CreateHotelSqlStore
	mongoStore CreateHotelMongoStore
}

func NewCreateHotelRepo(sqlStore CreateHotelSqlStore, mongoStore CreateHotelMongoStore) *createHotelRepo {
	return &createHotelRepo{
		sqlStore:   sqlStore,
		mongoStore: mongoStore,
	}
}

func (repo *createHotelRepo) Create(ctx context.Context, data *hotelmodel.HotelCreate) error {
	if err := repo.sqlStore.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(hotelmodel.EntityName, err)
	}

	if data.HotelAdditionalInfo == nil {
		return nil
	}

	data.HotelAdditionalInfo.HotelID = data.Id

	if err := repo.mongoStore.Create(ctx, data.HotelAdditionalInfo); err != nil {
		return common.ErrCannotCreateEntity(hotelmodel.EntityName, err)
	}

	return nil
}
