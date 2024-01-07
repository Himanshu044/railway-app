package lib

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectionPool struct {
	mutex          sync.Mutex
	maxConnections int
	connections    []*gorm.DB
}

var (
	conectionPool     *ConnectionPool
	conectionPoolOnce sync.Once
)

func InitializeConnectionPool(maxConnections int) *ConnectionPool {
	conectionPool := &ConnectionPool{
		maxConnections: maxConnections,
		connections:    make([]*gorm.DB, 0),
	}
	return conectionPool
}

func GetConnectionPool() *ConnectionPool {
	if conectionPool == nil {
		return InitializeConnectionPool(5)
	}
	return conectionPool
}

func (pool *ConnectionPool) GetConnection() *gorm.DB {
	if len(pool.connections) == 0 {
		db := createConnection()
		pool.connections = append(pool.connections, db)
		return db
	}
	conn := pool.connections[len(pool.connections)-1]
	pool.connections = pool.connections[:len(pool.connections)-1]
	return conn
}

func createConnection() *gorm.DB {
	databaseUrl := os.Getenv("Database_Connection_Url")

	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Error:", err)
	}
	return db
}

func (pool *ConnectionPool) ReleaseConnection(db *gorm.DB) {
	if len(pool.connections) >= pool.maxConnections {
		sqlDB, _ := db.DB()
		sqlDB.Close()
		return
	}
	pool.connections = append(pool.connections, db)
}
