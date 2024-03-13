package workersqlstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	workermodel "h5travelotobackend/module/worker/model"
)

func (s *sqlStore) FindWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*workermodel.Worker, error) {

	var data workermodel.Worker

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
