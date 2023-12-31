package helpers

import (
	"backend/lib"

	"gorm.io/gorm"
)

func GetDbInstance() *gorm.DB {
	pool := lib.GetConnectionPool()
	return pool.GetConnection()
}
