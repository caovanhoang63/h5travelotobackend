package htsavestore

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

func (s *store) FindSavedHotel(ctx context.Context, conditions map[string]interface{}) (*htsavemodel.HotelSave, error) {
	var result htsavemodel.HotelSave
	if err := s.db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
