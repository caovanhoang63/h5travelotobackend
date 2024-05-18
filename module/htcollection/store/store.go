package htcollectionstore

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
	es *elasticsearch.TypedClient
}

func NewStore(db *gorm.DB) *store {
	return &store{db: db}
}
