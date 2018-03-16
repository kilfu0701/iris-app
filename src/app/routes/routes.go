package routes

import (
	"github.com/kataras/iris"
)

func Init(app *iris.Application) {
	AppRoutes(app)
	ApiRoutes(app)
	//adminRoutes(app)
}
