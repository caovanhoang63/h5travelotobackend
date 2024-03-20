package workerbiz

import (
	"context"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

type DeleteWorkerStore interface {
	Delete(ctx context.Context, userId int) error
}

type deleteWorkerBiz struct {
	store     DeleteWorkerStore
	findStore FindWorkerStore
}

func NewDeleteWorkerBiz(store DeleteWorkerStore, findStore FindWorkerStore) *deleteWorkerBiz {
	return &deleteWorkerBiz{store: store, findStore: findStore}
}

func (biz *deleteWorkerBiz) DeleteWorker(ctx context.Context, hotelId int, userId int) error {
	_, err := biz.findStore.FindWithCondition(ctx, map[string]interface{}{"user_id": userId, "hotel_id": hotelId})
	if err != nil {
		return workermodel.ErrWorkerNotExist
	}

	if err := biz.store.Delete(ctx, userId); err != nil {
		return common.ErrCannotDeleteEntity(workermodel.EntityName, err)
	}
	return nil
}
