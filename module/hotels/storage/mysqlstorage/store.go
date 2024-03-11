package hotelmysqlstorage

import (
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

//
//func (s *sqlStore) BeginTransaction(ctx context.Context) {
//	s.db = s.db.WithContext(ctx).Begin()
//}
//
//func (s *sqlStore) RollBack(ctx context.Context) {
//	s.db.WithContext(ctx).Rollback()
//}
//
//func (s *sqlStore) Commit(ctx context.Context) {
//	s.db.WithContext(ctx).Commit()
//}
