package htcollectionstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

func (s *store) Update(ctx context.Context, id int,
	update *htcollection.HotelCollectionUpdate,
) error {
	if err := s.db.Table(htcollection.HotelCollection{}.TableName()).
		Where("id = ?", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
