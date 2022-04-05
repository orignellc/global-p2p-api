package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/app/logics"
	"github.com/orignellc/global-p2p-api/database"
	"github.com/orignellc/global-p2p-api/pkg/util"
	"strconv"
)


func GetClients(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(logics.Clients())
}

func GetClientByID(ctx *fiber.Ctx) error {
	strIdVar := ctx.Params("id")
	intIdVar, _ := strconv.Atoi(strIdVar)

	_client := logics.GetClientById(uint(intIdVar))
	_httpStatus := fiber.StatusOK

	if  _client.ID == 0 {
		_httpStatus = fiber.StatusNotFound
	}

	return ctx.Status(_httpStatus).JSON(_client)
}


func AddNewClient(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.HeaderContentType,fiber.MIMEApplicationJSON)

	var clientData = new(database.ClientData)

	if err := ctx.BodyParser(&clientData); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, errorWrapper := logics.NewClient(clientData)

	if errorWrapper != nil {
		switch errorWrapper.Type {
		case util.ValidationError:
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errorWrapper.Details)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(data.Public())
}

func UpdateClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}

func DeleteClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}