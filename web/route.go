package web

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) *fiber.App {
	api := app.Group("/api", Api)

	container := api.Group("/containers", Container)
	container.Get("/", GetAllContainer)
	container.Get("/:id", GetByIdContainer)
	container.Get("/logs/:id", GetByIdContainerLog)
	container.Post("/pause/:id", PauseContainer)
	container.Post("/unPause/:id", UnPauseContainer)
	container.Post("/start/:id", StartContainer)
	container.Post("/restart/:id", RestartContainer)
	container.Post("/stop/:id", StopContainer)
	container.Post("/kill/:id", KillContainer)
	container.Post("/rename/:id", RenameContainer)
	container.Post("/remove/:id", RemoveContainer)
	container.Post("/prune/", PruneContainer)
	// container.Post("/update/:id", UpdateContainer)

	api.Get("/network/:id")
	return app
}
