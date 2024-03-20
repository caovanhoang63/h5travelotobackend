package workersqlstorage

import (
	"context"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

func (s *sqlStore) Create(ctx context.Context, worker *workermodel.WorkerCreate) error {
	if err := s.db.Create(worker).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
