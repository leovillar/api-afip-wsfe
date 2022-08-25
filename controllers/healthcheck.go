package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leovillar/api-afip-wsfe/infra"
)

func Healthcheck(c *fiber.Ctx) error {

	if err := infra.DB.Ping(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return FEDummy(c)

}
