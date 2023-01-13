package auth

import (
	_ "backend.go/docs"
	"backend.go/src/controllers/authcontroller"
	_ "backend.go/src/models/authModel"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Login
// @Tags         Auth
// @Produce      json
// @Param        auth body authModel.Login true "Use response token to access other requests"
// @Success      200  {object}  authModel.Login
// @Failure      400
// @Failure      500
// @Router       /login [post]
func Login(app *fiber.App) {
	app.Post("/login", authcontroller.Login)
}

// @Summary      Register
// @Tags         Auth
// @Produce      json
// @Param        auth body authModel.Register true "Create new user"
// @Success      200  {object}  authModel.Register
// @Failure      400
// @Failure      500
// @Router       /register [post]
func Register(app *fiber.App) {
	app.Post("/register", authcontroller.Register)
}
