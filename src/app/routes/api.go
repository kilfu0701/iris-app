package routes

import (
	"github.com/kataras/iris"

	"app/controllers/api"
)

func ApiRoutes(app *iris.Application) {
	apiRoutes := app.Party("/api", loginMiddleware)

	apiRoutes.Get("/hello", api.Hello)
}

func loginMiddleware(ctx iris.Context) {
	ctx.Next()
}
