package routes

import (
	_ "backend.go/docs"
	"backend.go/src/middleware/auth"
	authentication "backend.go/src/routes/auth"
	"backend.go/src/routes/post"
	"backend.go/src/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Routes(app *fiber.App) {

	app.Group("/api", auth.IsAuth)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		err := c.SendString("Hello, World 👋!")
		return err
	})

	authentication.Login(app)

	authentication.Register(app)

	user.UserRoutes(app.Group("/api/v1"))

	post.PostRoutes(app.Group("/api/v1"))
}
