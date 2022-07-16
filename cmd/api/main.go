package main

import (
	"github.com/hmdrzaa11/got/pkg/app"
	"github.com/hmdrzaa11/got/pkg/routes"
)

func main() {
	app := app.Boot()

	//load routes
	routes.LoadRoutes(app)
	go func() {
		app.Run()
	}()

	app.GracefullShutdown()
}
