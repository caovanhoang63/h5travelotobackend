package pestore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	pemodel "h5travelotobackend/payment/module/paymentevent/model"
)

func (s *store) UpdateStatusToSuccess(ctx context.Context, txnId string) error {
	if err := s.db.WithContext(ctx).Table(pemodel.PaymentEvent{}.TableName()).
		Where("txn_id = ?", txnId).
		Update("is_payment_done", true).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
