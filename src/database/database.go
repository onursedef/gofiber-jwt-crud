package database

import (
	"sync"

	"backend.go/src/database/initialization"
	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initialization.Init()
	})

	return connection
}
