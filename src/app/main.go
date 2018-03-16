package main

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"app/config"
	"app/routes"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// -------------
	// load env.yaml
	// -------------
	config, err := config.Load(config.CONFIG_FILE)
	if err != nil {
		app.Logger().Fatalf("%++v", err)
	}

	// -------------
	// init routes
	// -------------
	routes.Init(app)

	// Register a view engine on .html files inside the ./view/** directory.
	app.RegisterView(iris.HTML("./views", ".html"))

	// http://localhost:9000
	app.Run(iris.Addr(":" + strconv.Itoa(config.Env.Port)), iris.WithoutServerError(iris.ErrServerClosed))
}
