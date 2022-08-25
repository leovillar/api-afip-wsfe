package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/leovillar/api-afip-wsfe/controllers"
)

func APIRoutes(app *fiber.App) {

	api := app.Group("/afip-wsfe/api")
	v1 := api.Group("/v1")

	v1.Get("/dashboard", monitor.New())
	v1.Get("/healthcheck", controllers.Healthcheck)
	v1.Get("/FEDummy", controllers.FEDummy)
	v1.Post("/LoginCMS", controllers.LoginCMS)
	v1.Post("/FECompUltimoAutorizado", controllers.FECompUltimoAutorizado)
	v1.Post("/FECAESolicitar", controllers.FECAESolicitar)
	v1.Get("/FEParamGetPtosVenta", controllers.FEParamGetPtosVenta)
	v1.Post("/FECompConsultar", controllers.FECompConsultar)

}
