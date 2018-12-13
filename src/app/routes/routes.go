package routes

import (
	"github.com/kataras/iris"

	"app/routes/user"
	//"src/app/routes/admin"
)

func Init(app *iris.Application) {
	user_routes.AppRoutes(app)
	user_routes.ApiRoutes(app)
	//adminRoutes(app)
}
