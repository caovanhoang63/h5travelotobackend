package userstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

func (s *sqlStore) ListUsersByIds(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	db := s.db.Table(common.SimpleUser{}.TableName())
	var users []common.SimpleUser
	if err := db.Where("id IN (?)", ids).Find(&users).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return users, nil
}
