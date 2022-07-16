package routes

import (
	"github.com/hmdrzaa11/got/pkg/app"
	"github.com/hmdrzaa11/got/pkg/handlers"
	"github.com/hmdrzaa11/got/pkg/middleware"
	"github.com/hmdrzaa11/got/pkg/models"
	"github.com/hmdrzaa11/got/pkg/services"
)

// LoadRoutes is going to register all the routes and handlers to the mux
func LoadRoutes(app *app.Application) {
	ps := services.NewDefaultProductsService(models.NewProductRepository())
	ph := handlers.NewProductHandler(app, ps)
	chain := middleware.CreateChain(middleware.DummyMiddleware, middleware.RequestLoggerMiddleware).Then(ph)
	app.Router.Handle("/", chain)
}
