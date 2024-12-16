package web

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) *fiber.App {
	app.Get("/containers", GetAllContainer)      // get All Container
	app.Get("/containers/:id", GetByIdContainer) // get by container id
	app.Get("/network/:id")
	return app
}
