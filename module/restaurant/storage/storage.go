package restaurantstorage

import "gorm.io/gorm"

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

type sqlStore struct {
	db *gorm.DB
}
