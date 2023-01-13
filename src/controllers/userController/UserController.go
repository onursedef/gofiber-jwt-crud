package usercontroller

import (
	"fmt"

	"backend.go/src/database"
	"backend.go/src/models/general"
	"backend.go/src/models/user"
	userRepository "backend.go/src/repository/userRepository"
	"github.com/gofiber/fiber/v2"
)

func Get(ctx *fiber.Ctx) error {

	users, err := userRepository.GetAll()

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to fetch users from the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Users",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	if len(users) == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: ctx.OriginalURL(),
			Title:       "No users found",
			Detail:      "There is no users found in the database. Try adding one and try again.",
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "Users",
			Error:   ErrorResponse,
		}

		return ctx.Status(404).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "Users",
		Message: fmt.Sprintf("Found %d users", len(users)),
		Data:    users,
	}

	return ctx.Status(200).JSON(response)
}

func Find(c *fiber.Ctx) error {

	user, err := userRepository.GetOne(c.Params("id"))

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to fetch user from the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(404).JSON(response)
	}

	if user.Id == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "No user found",
			Detail:      "There is no user found in the database. Try adding one and try again.",
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(404).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "User",
		Message: fmt.Sprintf("Found user with id %d", user.Id),
		Data:    user,
	}

	return c.Status(200).JSON(response)
}

func Create(c *fiber.Ctx) error {

	user := user.User{}

	if err := c.BodyParser(&user); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Request",
			Detail:      "The request you sent is invalid. Please check the request body and try again.",
		}

		response := general.Response{
			Status:  400,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(400).JSON(response)
	}

	user, err := userRepository.Create(user)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to create user in the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  201,
		Success: true,
		Kind:    "User",
		Message: fmt.Sprintf("Created user with id %d", user.Id),
		Data:    user,
	}

	return c.Status(201).JSON(response)
}

func Update(c *fiber.Ctx) error {

	user := user.User{}

	if err := c.BodyParser(&user); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Request",
			Detail:      "The request you sent is invalid. Please check the request body and try again.",
		}

		response := general.Response{
			Status:  400,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(400).JSON(response)
	}

	user, err := userRepository.Update(user)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to update user in the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "User",
		Message: fmt.Sprintf("Updated user with id %d", user.Id),
		Data:    user,
	}

	return c.Status(200).JSON(response)
}

func Patch(c *fiber.Ctx) error {

	user := user.User{}

	if err := c.BodyParser(&user); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Request",
			Detail:      "The request you sent is invalid. Please check the request body and try again.",
		}

		response := general.Response{
			Status:  400,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(400).JSON(response)
	}

	user, err := userRepository.Patch(c.Params("id"), user)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to patch user in the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "User",
		Message: fmt.Sprintf("Patched user with id %d", user.Id),
		Data:    user,
	}

	return c.Status(200).JSON(response)
}

func Delete(c *fiber.Ctx) error {
	user := user.User{}

	database.Connection().First(&user, c.Params("id"))

	if user.Id == 0 {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Not Found",
			Detail:      "The requested resource was not found.",
		}

		response := general.Response{
			Status:  404,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(404).JSON(response)
	}

	user, err := userRepository.Delete(user)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "A Server Error Occured",
			Detail:      fmt.Sprintf("An error occured while trying to delete user from the database. Error: %s", err),
		}

		response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "User",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(response)
	}

	response := general.Response{
		Status:  200,
		Success: true,
		Kind:    "User",
		Message: fmt.Sprintf("Deleted user with id %d", user.Id),
		Data:    user,
	}

	return c.Status(200).JSON(response)
}
