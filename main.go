package main

import (
	"backend.go/src/database/migration"
	"backend.go/src/utils/config"
	"backend.go/src/utils/server"
	"backend.go/src/utils/swagger"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

// @title Go Fiber CRUD API
// @description This is a sample server for a Go Fiber CRUD API.
// @version 1.0
// @schemes http https
// @host localhost:3001
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if !fiber.IsChild() {
		migration.Migrate()
	}

	config.LoadEnv()
	swagger.Configure()
	server.Start()
}
