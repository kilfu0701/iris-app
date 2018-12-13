package user_routes

import (
	"github.com/kataras/iris"

	"app/controllers/api"
)

func ApiRoutes(app *iris.Application) {
	apiRoutes := app.Party("/api", loginMiddleware)

	apiRoutes.Get("/hello", api.Hello)
	apiRoutes.Get("/user/{id:int min(1)}", api.GetUser)
}

func loginMiddleware(ctx iris.Context) {
	ctx.Next()
}
