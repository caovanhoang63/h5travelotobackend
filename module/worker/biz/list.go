package workerbiz

import (
	"context"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

type ListWorkerStore interface {
	ListHotelWorkers(ctx context.Context,
		conditions map[string]interface{},
		filter *workermodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.SimpleUser, error)
}

type listWorkerBiz struct {
	store ListWorkerStore
}

func NewListWorkerBiz(store ListWorkerStore) *listWorkerBiz {
	return &listWorkerBiz{store: store}
}

func (biz *listWorkerBiz) GetHotelWorkers(ctx context.Context,
	filter *workermodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {

	users, err := biz.store.ListHotelWorkers(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(workermodel.EntityName, err)
	}
	return users, nil
}
