package walletbiz

import (
	"golang.org/x/net/context"
)

type WithdrawalStore interface {
	Withdrawal(ctx context.Context, id int, amount float64) error
}

type withdrawalBiz struct {
	store WithdrawalStore
}

func NewWithdrawalBiz(store WithdrawalStore) *withdrawalBiz {
	return &withdrawalBiz{store: store}
}

func (w *withdrawalBiz) Withdrawal(ctx context.Context, id int, amount float64) error {
	err := w.store.Withdrawal(ctx, id, amount)

	if err != nil {
		return err
	}

	return nil
}
