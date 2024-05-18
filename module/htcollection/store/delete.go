package htcollectionstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/htcollection/model"
)

func (s *store) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(htsavemodel.HotelCollection{}.TableName()).
		Where("id = ?", id).
		Update("status", common.StatusDeleted).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *store) RemoveHotelFromCollection(ctx context.Context, hotelId, collectionId int) error {
	if err := s.db.Table(htsavemodel.HotelCollectionDetail{}.TableName()).
		Where("hotel_id = ? AND collection_id = ?", hotelId, collectionId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
