package dealsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

func (s *sqlStore) DeleteDeal(ctx context.Context, id int) error {
	if err := s.db.Table(dealmodel.Deal{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
