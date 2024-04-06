package hoteldetailsqlstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
