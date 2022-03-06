package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/app/models"
	"github.com/orignellc/global-p2p-api/database"
	"github.com/orignellc/global-p2p-api/pkg/util"
	"gorm.io/gorm"
)

type ClientData struct {
	models.ClientBasic
	DomainWhitelists []string `json:"domain_whitelists"`
}

var db *gorm.DB

func init() {
	db = database.GetDB()
}

func (*ClientData) ModelName() string {
	return "client"
}

func GetClients(ctx *fiber.Ctx) error {
	var clients []ClientData

	db.Find(clients)

	return ctx.Status(fiber.StatusOK).JSON(clients)
}

func GetClientByID(ctx *fiber.Ctx) error {

	return ctx.JSON(ctx.App().Stack())
}

func AddNewClient(ctx *fiber.Ctx) error {
	ctx.Accepts(fiber.HeaderContentType,fiber.MIMEApplicationJSON)

	var client ClientData

	if err := ctx.BodyParser(&client); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := util.ValidateStruct(&client)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	db.Create(client)

	return ctx.Status(fiber.StatusOK).JSON(client)
}

func UpdateClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}

func DeleteClientByID(ctx *fiber.Ctx) error {
	return ctx.JSON(ctx.App().Stack())
}