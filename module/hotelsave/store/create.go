package htsavestore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

func (s *store) Create(ctx context.Context, data *htsavemodel.HotelSaveCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
