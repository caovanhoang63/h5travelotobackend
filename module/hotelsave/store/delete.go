package htsavestore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

func (s *store) Delete(ctx context.Context, del *htsavemodel.HotelSaveDelete) error {
	if err := s.db.Table(del.TableName()).Where(del).Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
