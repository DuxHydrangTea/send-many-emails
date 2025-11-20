package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-proj/mail"
	// "fmt"
)

func Register(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return c.Status(500).JSON(fiber.Map{"error" : "Email is null now!"})
	}

	go mail.SendEmail(
		email,
		"Welcome to my system!",
		"<h2>Hello con cho</h2>",
	)

	return c.JSON(fiber.Map{
		"message": "Send successfully",
		"email": email,
	})
}

func SendManyEmails(c *fiber.Ctx) error {
	emails := []string{}

	for i := 1; i <= 40; i++{
		emails = append(emails, "ngocdung2002d@gmail.com")
	}

	mail.SendMassEmail(emails)
	
	return c.Status(200).JSON(fiber.Map{
		"Message": "Okeluon",
	})
}