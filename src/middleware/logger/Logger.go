package logger

import (
	"log"
	"os"
	"path/filepath"

	"backend.go/src/utils/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) *os.File {
	basepath, _ := filepath.Abs("./logs")

	file, err := os.OpenFile(
		basepath+Filename(),               // /Users/x/go-backend-v1/logs/production.log
		os.O_RDWR|os.O_CREATE|os.O_APPEND, // Read, Write, Create, Append
		0666,                              // Permission
	)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "[${time}] | ${status} | ${ip}:${port} | ${latency} | ${method} | ${path}\n", // [15:04:05 02-Jan-2006] | 200 | 127.0.0.1:3001 | 14.6µs | GET | /api/v1/users
		TimeFormat: "15:04:05 02-Jan-2006",
		TimeZone:   "Europe/Istanbul",
	}))

	return file

}

func Filename() string {
	if config.GetAppEnv() == "production" {
		return "/production.log"
	} else {
		return "/development.log"
	}
}
