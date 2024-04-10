package dealsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

func (s *sqlStore) Create(ctx context.Context, deal *dealmodel.DealCreate) error {
	if err := s.db.Create(&deal).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
