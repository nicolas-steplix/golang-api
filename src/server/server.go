package server

import (
	"fmt"
	"golang-api/src/helpers"
	"golang-api/src/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartServer() error {
	port := helpers.Coalesce(viper.GetInt("PORT"), 8080)

	app := fiber.New()
	routes.RegisterRoutes(app)
	return app.Listen(fmt.Sprintf(":%d", port))
}
