package invoicestorage

import "gorm.io/gorm"

// InvoiceStorage defines the storage layer for invoice
type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}
