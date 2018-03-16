package api

import (
	"github.com/kataras/iris"
)

func Hello(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
