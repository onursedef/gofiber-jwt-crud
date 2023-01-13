package user

import (
	userController "backend.go/src/controllers/userController"
	_ "backend.go/src/models/user"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	GetAll(app.Group(""))
	GetOne(app.Group(""))
	Create(app.Group(""))
	Update(app.Group(""))
	Delete(app.Group(""))
}

// @Summary      Get all users
// @Tags         Users
// @Produce      json
// @Success      200  {object}  []user.User
// @Failure      404
// @Failure      500
// @Router       /api/v1/users [get]
// @Security     ApiKeyAuth
func GetAll(app fiber.Router) {
	app.Get("/users", userController.Get)
}

// @Summary      Get one user
// @Tags         Users
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200  {object}  user.User
// @Failure      404
// @Failure      400
// @Failure      500
// @Router       /api/v1/users/{id} [get]
// @Security     ApiKeyAuth
func GetOne(app fiber.Router) {
	app.Get("/users/:id", userController.Find)
}

// @Summary      Create a user
// @Tags         Users
// @Produce      json
// @Param        post body user.User true "Create a new user"
// @Success      201  {object}  user.User
// @Failure      400
// @Failure      500
// @Router       /api/v1/user [post]
// @Security     ApiKeyAuth
func Create(app fiber.Router) {
	app.Post("/user", userController.Create)
}

// @Summary      Update a user
// @Tags         Users
// @Produce      json
// @Param 	  	 id path int true "User ID"
// @Param     	 user body user.User true "Update a user"
// @Success   	 200  {object}  user.User
// @Failure      404
// @Failure      400
// @Failure      500
// @Router    	/api/v1/user/{id} [put]
// @Security     ApiKeyAuth
func Update(app fiber.Router) {
	app.Put("/user/:id", userController.Update)
}

// @Summary      Patch a user
// @Tags         Users
// @Produce      json
// @Param        id path int true "User ID"
// @Param        user body user.User true "Patch a User"
// @Success  	 200 {object} user.User
// @Failure      404
// @Failure      400
// @Failure      500
// @Router   	/api/v1/user/{id} [patch]
// @Security     ApiKeyAuth
func Patch(app fiber.Router) {
	app.Patch("/user/:id", userController.Patch)
}

// @Summary      Delete a user
// @Tags         Users
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200  {object}  user.User
// @Failure      404
// @Failure      400
// @Failure      500
// @Router       /api/v1/user/{id} [delete]
// @Security     ApiKeyAuth
func Delete(app fiber.Router) {
	app.Delete("/user/:id", userController.Delete)
}
