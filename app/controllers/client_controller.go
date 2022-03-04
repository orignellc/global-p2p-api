package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/app/models"
)

type ClientData struct {
	models.ClientBasic
	DomainWhitelists []string `json:"domain_whitelists"`
}

func GetClients(ctx *fiber.Ctx) error {
	//ctx.Accepts(fiber.HeaderContentType,fiber.MIMEApplicationJSON)

	var client ClientData

	if err := ctx.BodyParser(&client); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(client)
}

func GetClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}

func AddNewClient(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}

func UpdateClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}

func DeleteClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}