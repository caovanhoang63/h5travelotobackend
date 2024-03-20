package workerbiz

import (
	"context"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

type CreateWorkerStore interface {
	Create(ctx context.Context, worker *workermodel.WorkerCreate) error
}

type createWorkerBiz struct {
	store     CreateWorkerStore
	findStore FindWorkerStore
}

func NewCreateWorkerBiz(store CreateWorkerStore, findStore FindWorkerStore) *createWorkerBiz {
	return &createWorkerBiz{store: store, findStore: findStore}
}

func (biz *createWorkerBiz) CreateWorker(ctx context.Context, worker *workermodel.WorkerCreate) error {
	if _, err := biz.findStore.FindWithCondition(ctx, map[string]interface{}{"user_id": worker.UserID}); err == nil {
		return workermodel.ErrWorkerAlreadyExist
	}

	if err := biz.store.Create(ctx, worker); err != nil {
		return common.ErrCannotCreateEntity(workermodel.EntityName, err)
	}
	return nil
}
