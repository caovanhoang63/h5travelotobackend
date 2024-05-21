package payinstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

func (s *store) Create(ctx context.Context,
	create *payinmodel.PaymentBookingCreate) error {
	if err := s.db.WithContext(ctx).Create(create).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
