package dealsqlstorage

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

func (s *sqlStore) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*dealmodel.Deal, error) {
	var deal dealmodel.Deal
	if err := s.db.Where(condition).First(&deal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &deal, nil
}
