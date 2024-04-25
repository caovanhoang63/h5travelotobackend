package invoicebiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	invoicemodel "h5travelotobackend/payment/module/invoices/model"
)

type CreateInvoiceStore interface {
	Create(ctx context.Context, data *invoicemodel.InvoiceCreate) error
}

type createInvoiceBiz struct {
	store CreateInvoiceStore
}

func NewCreateInvoiceBiz(store CreateInvoiceStore) *createInvoiceBiz {
	return &createInvoiceBiz{store: store}
}

func (biz *createInvoiceBiz) CreateInvoice(ctx context.Context, data *invoicemodel.InvoiceCreate) error {

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
