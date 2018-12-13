package api

import (
	"github.com/kataras/iris"

	"app/models/user"
)

func GetUser(ctx iris.Context) {
	userID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}


	user := user.User{ID: 1}
	if ok, _ := orm.Get(&user); ok {
		ctx.Writef("user found: %#v", user)
		return
	}


	ctx.JSON(map[string]interface{}{
		"userID": userID,
		"user": user,
	})
}
