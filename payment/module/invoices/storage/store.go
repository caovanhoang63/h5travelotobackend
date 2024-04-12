package storage

import "gorm.io/gorm"

// InvoiceStorage defines the storage layer for invoice
type sqlStore struct {
	gorm *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}
