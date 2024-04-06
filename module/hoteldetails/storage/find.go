package hoteldetailsqlstorage

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

func (s *sqlStore) FindWithCondition(ctx context.Context, condition map[string]interface{}) (*hoteldetailmodel.HotelDetail, error) {
	var data hoteldetailmodel.HotelDetail

	if err := s.db.First(&data, condition).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &data, nil
}
