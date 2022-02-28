package controllers

import "github.com/gofiber/fiber/v2"

func GetClients(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
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