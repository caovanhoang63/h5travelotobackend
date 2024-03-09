package userstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	usermodel "h5travelotobackend/module/users/model"
)

func (s *sqlStore) FindUser(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &user, nil
}
