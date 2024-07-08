package walletstorage

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *store {
	return &store{db: db}
}
