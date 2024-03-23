package hotelrepo

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type FindHotelSqlStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hotelmodel.Hotel, error)
}

type FindHotelMongoStore interface {
	FindAdditionalInfo(ctx context.Context, id int) (*hotelmodel.HotelAdditionalInfo, error)
}

type findHotelRepo struct {
	sqlStore   FindHotelSqlStore
	mongoStore FindHotelMongoStore
}

func NewFindHotelRepo(sqlStore FindHotelSqlStore, mongoStore FindHotelMongoStore) *findHotelRepo {
	return &findHotelRepo{sqlStore: sqlStore, mongoStore: mongoStore}
}

func (r *findHotelRepo) FindBaseDataWithCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*hotelmodel.Hotel, error) {
	result, err := r.sqlStore.FindDataWithCondition(ctx, condition, moreKeys...)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}

	return result, nil
}

func (r *findHotelRepo) FindAdditionalDataById(ctx context.Context, id int) (*hotelmodel.HotelAdditionalInfo, error) {
	data, err := r.mongoStore.FindAdditionalInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *findHotelRepo) FindAllDataWithCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*hotelmodel.Hotel, error) {
	data, err := r.FindBaseDataWithCondition(ctx, condition, moreKeys...)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}
	//data.HotelAdditionalInfo, err = r.mongoStore.FindAdditionalInfo(ctx, data.Id)

	if err != nil {
		return nil, err
	}
	return data, nil

}
