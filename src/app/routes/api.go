package routes

import (
	"github.com/kataras/iris"

	"app/controllers/api"
)

func ApiRoutes(app *iris.Application) {
	apiRoutes := app.Party("/api", apiMiddleware)

	apiRoutes.Get("/hello", api.Hello)
}

func apiMiddleware(ctx iris.Context) {
	ctx.Next()
}
