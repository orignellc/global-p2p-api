package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/orignellc/global-p2p-api/app/controllers"
)

type ApiV1Router struct {
}

func (h ApiV1Router) InstallRouter(app *fiber.App) {

	api := app.Group("/api/v1/", limiter.New())

	api.Route("clients/", func(api fiber.Router) {
		// List clients
		// GET /api/v1/clients/all
		api.Get("all", controllers.GetClients).Name("all")

		// Get client by ID
		// GET /api/v1/clients/1
		api.Get(":id", controllers.GetClientByID).Name("show")

		// Add a new client
		// POST /api/v1/clients/create
		api.Post("create",controllers.AddNewClient).Name("create")

		// Update a client by ID
		// Patch /api/v1/clients/update/1
		api.Patch("update/:id", controllers.UpdateClientByID).Name("update")

		// Delete client by ID
		// Delete /api/v1/clients/delete/1
		api.Delete("delete/:id", controllers.DeleteClientByID).Name("delete")

	},"clients.")

	api.Route("countries/", func(api fiber.Router) {
	// List countries
	// GET /api/v1/countries/all
	api.Get("all", controllers.GetCountries).Name("all")

	// Get client by ID
	// GET /api/v1/countries/1
	api.Get(":id", controllers.GetCountryByID).Name("show")

	},"countries.")
}

func NewApiV1Router() *ApiV1Router {
	return &ApiV1Router{}
}