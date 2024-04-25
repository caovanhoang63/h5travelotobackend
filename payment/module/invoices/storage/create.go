package invoicestorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	invoicemodel "h5travelotobackend/payment/module/invoices/model"
)

func (s *sqlStore) Create(ctx context.Context, data *invoicemodel.InvoiceCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
