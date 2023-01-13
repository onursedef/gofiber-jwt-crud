package post

import (
	"backend.go/src/controllers/postcontroller"
	_ "backend.go/src/models/post"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app fiber.Router) {
	GetAll(app.Group(""))
	GetOne(app.Group(""))
	Create(app.Group(""))
	Update(app.Group(""))
	Delete(app.Group(""))
}

// @Summary      Get all posts
// @Tags         Posts
// @Produce      json
// @Success      200  {object}  []post.Post
// @Failure      404
// @Failure      500
// @Router       /api/v1/posts [get]
// @Security     ApiKeyAuth
func GetAll(app fiber.Router) {
	app.Get("/posts", postcontroller.Get)
}

// @Summary      Get one post
// @Tags         Posts
// @Produce      json
// @Param        id path int true "Post ID"
// @Success      200  {object}  post.Post
// @Failure      404
// @Failure      400
// @Failure      500
// @Router       /api/v1/posts/{id} [get]
// @Security     ApiKeyAuth
func GetOne(app fiber.Router) {
	app.Get("/posts/:id", postcontroller.Find)
}

// @Summary      Create a post
// @Tags         Posts
// @Produce      json
// @Param        post body post.Post true "Create a new post"
// @Success      201  {object}  post.Post
// @Failure      400
// @Failure      500
// @Router       /api/v1/post [post]
// @Security     ApiKeyAuth
func Create(app fiber.Router) {
	app.Post("/post", postcontroller.Create)
}

// @Summary      Update a post
// @Tags         Posts
// @Produce      json
// @Param 	  	 id path int true "Post ID"
// @Param     	 post body post.Post true "Update a post"
// @Success   	 200  {object}  post.Post
// @Failure      404
// @Failure      400
// @Failure      500
// @Router    	/api/v1/post/{id} [put]
// @Security     ApiKeyAuth
func Update(app fiber.Router) {
	app.Put("/post/:id", postcontroller.Update)
}

// @Summary      Patch a post
// @Tags         Posts
// @Produce      json
// @Param        id path int true "Post ID"
// @Param        post body post.Post true "Patch a post"
// @Success  	 200 {object} post.Post
// @Failure      404
// @Failure      400
// @Failure      500
// @Router   	/api/v1/post/{id} [patch]
// @Security     ApiKeyAuth
func Patch(app fiber.Router) {
	app.Patch("/post/:id", postcontroller.Patch)
}

// @Summary      Delete a post
// @Tags         Posts
// @Produce      json
// @Param        id path int true "Post ID"
// @Success      200  {object}  post.Post
// @Failure      404
// @Failure      400
// @Failure      500
// @Router       /api/v1/post/{id} [delete]
// @Security     ApiKeyAuth
func Delete(app fiber.Router) {
	app.Delete("/post/:id", postcontroller.Delete)
}
