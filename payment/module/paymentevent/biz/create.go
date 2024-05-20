package pebiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/uuid"
	pemodel "h5travelotobackend/payment/module/paymentevent/model"
)

type PECreateStore interface {
	Create(ctx context.Context, create *pemodel.PaymentEventCreate) error
}

type PEBiz struct {
	peCreateStore PECreateStore
	uuid          uuid.Uuid
}

func NewPEBiz(peCreateStore PECreateStore, uuid uuid.Uuid) *PEBiz {
	return &PEBiz{peCreateStore: peCreateStore, uuid: uuid}
}

func (biz *PEBiz) Create(ctx context.Context, create *pemodel.PaymentEventCreate) error {
	if txnId, err := biz.uuid.Generate(); err != nil {
		return common.ErrInternal(err)
	} else {
		create.TxnId = txnId
	}
	if err := biz.peCreateStore.Create(ctx, create); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
