package database

import (
	"sync"

	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initializePostgres()
	})

	return connection
}
