package payinstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	payinmodel "h5travelotobackend/payment/module/payin/model"
)

func (s *store) UpdateStatus(ctx context.Context, txnId string, status *payinmodel.PaymentBookingUpdateStatus) error {
	if err := s.db.WithContext(ctx).Where("txn_id = ?", txnId).Updates(status).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
