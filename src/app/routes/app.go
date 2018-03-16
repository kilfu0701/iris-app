package routes

import (
	"github.com/kataras/iris"

	"app/controllers"
)

func AppRoutes(app *iris.Application) {
	app.Get("/", controllers.Top)
}
