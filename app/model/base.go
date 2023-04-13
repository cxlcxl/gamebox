package model

import "gorm.io/gorm"

var (
	cacheDbKey = "db"
)

type db struct {
	*gorm.DB
}
