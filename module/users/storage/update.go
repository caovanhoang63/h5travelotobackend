package userstorage

import (
	"context"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *usermodel.UserUpdate) error {
	if err := s.db.Where("id = ? ", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) ChangePassword(ctx context.Context, id int, data *usermodel.UserChangePassword) error {
	if err := s.db.Where("id = ? ", id).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil

}
