package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	middleware "backend.go/src/middleware/logger"
	"backend.go/src/routes"
	"backend.go/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: utils.ErrorHandler,
	})

	file := middleware.Logger(app)

	app.Use(
		logger.New(),
	)

	defer file.Close()

	routes.Routes(app)

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
