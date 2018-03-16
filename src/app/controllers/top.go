package controllers

import (
	"github.com/kataras/iris"
)

func Top(ctx iris.Context) {
	ctx.Application().Logger().Infof("Top")
	ctx.HTML("<h1>Welcome</h1>")
}
