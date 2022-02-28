package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/orignellc/global-p2p-api/app/controllers"
)

type ApiV1Router struct {
}

func (h ApiV1Router) InstallRouter(app *fiber.App) {
	api := app.Group("/api/v1", limiter.New())

	api.Route("clients/", func(api fiber.Router) {
		// List clients
		api.Get("all", controllers.GetClients).Name("all")

		// Get client by ID
		api.Get("/{id:[0-9]+}", controllers.GetClientByID).Name("show")

		// Add a new client
		api.Post("/create",controllers.AddNewClient).Name("create")

		// Update a client by ID
		api.Patch("/update/{id:[0-9]+}", controllers.UpdateClientByID).Name("update")

		// Delete client by ID
		api.Delete("/delete/{id:[0-9]+}", controllers.DeleteClientByID).Name("delete")
	},"clients.")
}

func NewApiV1Router() *ApiV1Router {
	return &ApiV1Router{}
}