package main

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

const (
	Admin = 1
	User  = 0
)

var UsersTokens = map[string]string{}

func Hello(c *fiber.Ctx) error {
	token := c.Query("token")
	_, exist := UsersTokens[token]
	if exist {
		err := c.SendString("Hello")
		if err != nil {
			return err
		}
	} else {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	return nil
}

func Registration(c *fiber.Ctx) error {
	user := struct {
		UserFirstName string `json:"user_first_name" validate:"required"`
		Email         string `json:"email" validate:"required"`
		UserRole      int64  `json:"user_role" validate:"required"`
	}{}

	err := c.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "can not parse body, check types of fields")
	}

	if user.UserRole != Admin && user.UserRole != User {
		return fiber.NewError(fiber.StatusBadRequest, "Not valid user role")
	}

	if len(user.UserFirstName) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Length of user first name must be > 0")
	}

	if len(user.Email) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Length of user email must be > 0")
	}

	token := generateToken(user.Email)

	UsersTokens[token] = user.Email

	err = c.SendString(token)
	if err != nil {
		return err
	}

	return nil
}

func generateToken(email string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
