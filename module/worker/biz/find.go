package workerbiz

import (
	"context"
	workermodel "h5travelotobackend/module/worker/model"
)

type FindWorkerStore interface {
	FindWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*workermodel.Worker, error)
}
