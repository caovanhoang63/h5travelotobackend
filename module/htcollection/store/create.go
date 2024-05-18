package htcollectionstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

func (s *store) Create(ctx context.Context, data *htcollection.HotelCollectionCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *store) AddHotelToCollection(ctx context.Context, data *htcollection.HotelCollectionDetailCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
