package postcontroller

import (
	"fmt"

	"backend.go/src/database"
	"backend.go/src/models/general"
	postrepository "backend.go/src/repository/postRepository"

	"backend.go/src/models/post"
	"github.com/gofiber/fiber/v2"
)

func Get(ctx *fiber.Ctx) error {
	posts, err := postrepository.GetAll()

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to fetch posts from the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	if len(posts) == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "No posts found",
			Detail:      "There is no post found in the database. Try adding one and try again.",
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "Posts",
		Message: fmt.Sprintf("Found %d posts", len(posts)),
		Data:    posts,
	}

	return ctx.Status(200).JSON(response)
}

func Find(ctx *fiber.Ctx) error {

	post, err := postrepository.GetOne(ctx.Params("id"))

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to fetch posts from the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(500).JSON(response)
	}

	if post.ID == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "No post found",
			Detail:      fmt.Sprintf("Sorry, but we were unable to locate a post in our database that matches the provided ID of %s. Please double-check the ID and try again.", ctx.Params("id")),
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "Posts",
		Message: fmt.Sprintf("Found post with ID of %s", ctx.Params("id")),
		Data:    post,
	}

	return ctx.Status(200).JSON(response)
}

func Create(ctx *fiber.Ctx) error {

	post := post.Post{}

	if err := ctx.BodyParser(&post); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "Service Unavailable",
			Detail:      fmt.Sprintf("An error occured while trying to parse the request body. Error: %s", err),
		}

		response := general.Response{
			Status:  503,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(503).JSON(response)
	}

	post, err := postrepository.Create(post)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to create a new post. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  201,
		Success: true,
		Message: fmt.Sprintf("Successfully created a new post with ID of %d", post.ID),
		Kind:    "Posts",
		Data:    post,
	}

	return ctx.Status(201).JSON(response)
}

// Will not be used but will be left here for reference
func Update(ctx *fiber.Ctx) error {
	post := post.Post{}

	if err := ctx.BodyParser(&post); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "Service Unavailable",
			Detail:      fmt.Sprintf("An error occured while trying to parse the request body. Error: %s", err),
		}

		response := general.Response{
			Status:  503,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(503).JSON(response)
	}

	post, err := postrepository.Update(post)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to update the post. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  201,
		Success: true,
		Message: fmt.Sprintf("Successfully updated post with ID of %d", post.ID),
		Kind:    "Posts",
		Data:    post,
	}

	return ctx.Status(201).JSON(response)
}

func Patch(ctx *fiber.Ctx) error {
	var post post.Post

	if err := ctx.BodyParser(&post); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "Service Unavailable",
			Detail:      fmt.Sprintf("An error occured while trying to parse the request body. Error: %s", err),
		}

		response := general.Response{
			Status:  503,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(503).JSON(response)
	}

	post, err := postrepository.Patch(ctx.Params("id"), post)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to patch the post. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  201,
		Success: true,
		Message: fmt.Sprintf("Successfully patched post with ID of %d", post.ID),
		Kind:    "Posts",
		Data:    post,
	}

	return ctx.Status(201).JSON(response)
}

func Delete(ctx *fiber.Ctx) error {
	post := post.Post{}

	database.Connection().First(&post, ctx.Params("id"))

	if post.ID == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "Not Found",
			Detail:      fmt.Sprintf("The post with ID of %s was not found", ctx.Params("id")),
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	post, err := postrepository.Delete(post)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to patch the post. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Posts",
			Error:   ErrorResponse,
		}

		return ctx.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  204,
		Success: true,
		Message: fmt.Sprintf("Successfully deleted post with ID of %d", post.ID),
		Kind:    "Posts",
		Data:    post,
	}

	return ctx.Status(204).JSON(response)
}
