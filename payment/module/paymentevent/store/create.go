package pestore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	pemodel "h5travelotobackend/payment/module/paymentevent/model"
)

func (s *store) Create(ctx context.Context, create *pemodel.PaymentEventCreate) error {
	if err := s.db.WithContext(ctx).Create(create).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
