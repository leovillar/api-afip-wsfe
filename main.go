package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leovillar/api-afip-wsfe/infra"
	"github.com/leovillar/api-afip-wsfe/routes"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func init() {
	gotenv.Load(".env")
}

func main() {

	infra.InitDB()
	defer infra.DB.Close()

	app := fiber.New()
	routes.APIRoutes(app)
	log.Fatal(app.Listen(":" + os.Getenv("API_PORT")))

}
