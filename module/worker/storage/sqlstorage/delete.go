package workersqlstorage

import (
	"context"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

func (s *sqlStore) Delete(ctx context.Context, userId int) error {
	if err := s.db.Where("user_id = ?", userId).Delete(&workermodel.Worker{}).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
