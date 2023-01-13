package initialization

import (
	"log"
	"os"

	"backend.go/src/database/drivers"
	"backend.go/src/utils/config"

	"gorm.io/gorm"
)

func Init() *gorm.DB {
	config.LoadEnv()

	switch os.Getenv("DB_DRIVER") {
	case "MySQL":
		return drivers.InitizalizeMysql()
	case "PostgreSQL":
		return drivers.InizitalizePostgres()
	default:
		log.Fatal("Database driver not found")
		return nil
	}
}
