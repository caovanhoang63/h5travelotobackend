package payinstore

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

func (s *store) FindWithCondition(ctx context.Context, cond map[string]interface{}) (*payinmodel.PaymentBooking, error) {
	var result payinmodel.PaymentBooking
	if err := s.db.WithContext(ctx).First(&result, cond).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}

func (s *store) FindExecutingOrSuccessByBookingId(ctx context.Context, bookingId int) (*payinmodel.PaymentBooking, error) {
	var result payinmodel.PaymentBooking
	if err := s.db.WithContext(ctx).
		Where("booking_id = ?", bookingId).
		Where("payment_status = 'success' OR" +
			" (payment_status = 'executing' AND TIMESTAMPDIFF(HOUR, created_at, NOW()) < 1 )").
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
