package dealsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *dealmodel.DealUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(&update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
