package main

import (
	"docker-debug/src/routes"

	iris "github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	setLogger(app)

	// routes
	setHelloRoutes(app)

	app.Run(iris.Addr(":5000"), iris.WithoutServerError(iris.ErrServerClosed))
}

func setHelloRoutes(app *iris.Application) {
	app.Get("/hello", routes.HelloWorld)
}

func setLogger(app *iris.Application) {
	app.Logger().SetLevel("info")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())
}
