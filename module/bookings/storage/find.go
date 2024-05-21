package bookingsqlstorage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) FindWithCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*bookingmodel.Booking, error) {
	var result bookingmodel.Booking

	if err := s.db.Where(condition).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	result.CheckExpired()

	return &result, nil
}

func (s *sqlStore) FindDTOWithCondition(ctx context.Context,
	condition map[string]interface{}) (*common.DTOBooking, error) {
	var result common.DTOBooking

	if err := s.db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
