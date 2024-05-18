package htcollectionstore

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

func (s *store) FindCollection(ctx context.Context,
	conditions map[string]interface{},
) (*htcollection.HotelCollection, error) {
	var result htcollection.HotelCollection
	if err := s.db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
