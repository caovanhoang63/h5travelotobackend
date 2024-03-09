package userstorage

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {

	db := s.db.Begin()

	if err := db.Table(usermodel.User{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}
	return nil
}
