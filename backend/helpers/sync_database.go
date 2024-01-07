package helpers

import (
	"backend/lib"
	"backend/models"
)

func SyncDatabase() {
	pool := lib.GetConnectionPool()
	DB := pool.GetConnection()
	DB.AutoMigrate(&models.User{})
}
