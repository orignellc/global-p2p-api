package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orignellc/global-p2p-api/app/logics"
	"strconv"
)


func GetCountries(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(logics.Countries())
}

func GetCountryByID(ctx *fiber.Ctx) error {
	strIdVar := ctx.Params("id")
	intIdVar, _ := strconv.Atoi(strIdVar)

	_country := logics.GetCountryById(uint(intIdVar))
	_httpStatus := fiber.StatusOK

	if  _country.ID == 0 {
		_httpStatus = fiber.StatusNotFound
	}

	return ctx.Status(_httpStatus).JSON(_country)
}