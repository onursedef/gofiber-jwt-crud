package authcontroller

import (
	"os"
	"time"

	"backend.go/src/models/authModel"
	"backend.go/src/models/general"
	"backend.go/src/models/user"
	userrepository "backend.go/src/repository/userRepository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data authModel.Login

	if err := c.BodyParser(&data); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Request",
			Detail:      "The request you sent is invalid. Please check the request body and try again.",
		}

		Response := general.Response{
			Status:  400,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(400).JSON(Response)
	}

	user, err := userrepository.GetByEmail(data.Email)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Credentials",
			Detail:      "The username or password you entered is incorrect.",
		}

		Response := general.Response{
			Status:  401,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(403).JSON(Response)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)

	if data.Email != user.Email && string(hashedPassword) != user.Password {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Credentials",
			Detail:      "The username or password you entered is incorrect.",
		}

		Response := general.Response{
			Status:  401,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(403).JSON(Response)
	}

	claims := jwt.MapClaims{
		"name":  "Administrator",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "An error occured",
			Detail:      "An error occured while trying to generate the token. Please try again.",
		}

		Response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(Response)
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}

func Register(c *fiber.Ctx) error {

	user := user.User{}

	if err := c.BodyParser(&user); err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "Invalid Request",
			Detail:      "The request you sent is invalid. Please check the request body and try again.",
		}

		Response := general.Response{
			Status:  400,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(400).JSON(Response)
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(hashedPassword)

	user, err := userrepository.Create(user)

	if err != nil {
		ErrorResponse := general.ErrorResponse{
			Request_url: c.OriginalURL(),
			Title:       "An error occured",
			Detail:      "An error occured while trying to create the user. Please try again.",
		}

		Response := general.Response{
			Status:  500,
			Success: false,
			Kind:    "Auth",
			Error:   ErrorResponse,
		}

		return c.Status(500).JSON(Response)
	}

	response := general.Response{
		Status:  201,
		Success: true,
		Kind:    "User",
		Data:    user,
		Message: "User registered successfully",
	}

	return c.Status(201).JSON(response)
}
