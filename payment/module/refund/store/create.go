package refundstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	refundmodel "h5travelotobackend/payment/module/refund/model"
)

func (s *store) Create(ctx context.Context,
	create *refundmodel.RefundBookingCreate) error {
	if err := s.db.WithContext(ctx).Create(create).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
